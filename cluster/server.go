package cluster

import (
	"fmt"
	"log"
	"multipass-cluster/agent"
	"net"
	"strings"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type server struct {
	UnimplementedRpcServer
	listener   net.Listener
	clients    *map[string]agent.Client
	grpcServer *grpc.Server
	clientsMu  sync.RWMutex
}

type Server interface {
	GetClients() *map[string]agent.Client
	Serve() error
}

func (s *server) GetClients() *map[string]agent.Client {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()
	return s.clients
}

func (s *server) addClient(uid string, client agent.Client) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()
	(*s.clients)[uid] = client
}

func (s *server) removeClient(uid string) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()
	err := (*s.clients)[uid].Close()
	if err != nil {
		log.Printf("failed to close agent client: %v", err)
	}
	delete(*s.clients, uid)
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) Join(stream grpc.BidiStreamingServer[JoinRequest, JoinReply]) error {
	id := uuid.Must(uuid.NewRandom()).String()
	defer log.Printf("client disconnected: %s", id)
	defer s.removeClient(id)

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("failed to get peer from context")
	}

	for {
		request, err := stream.Recv()
		if err != nil {
			return err
		}

		if err := stream.Send(&JoinReply{
			Uuid: id,
		}); err != nil {
			return fmt.Errorf("failed to send join reply on master: %w", err)
		}

		target := fmt.Sprintf("%s:%d", strings.Split(p.Addr.String(), ":")[0], request.AgentInfo.Port)
		agentClient, err := agent.NewClient(target)
		if err != nil {
			return fmt.Errorf("failed to create multipass client: %w", err)
		}

		s.addClient(id, agentClient)

		log.Printf("joined hostname: %s", request.Hostname)
	}
}

func NewServer(addr string) (Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := &server{
		clientsMu:  sync.RWMutex{},
		clients:    &map[string]agent.Client{},
		listener:   lis,
		grpcServer: grpcServer,
	}
	RegisterRpcServer(grpcServer, server)
	return server, nil
}
