package cluster

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"multipass-cluster/agent"
	"net"
	"strings"
	"sync"
)

type server struct {
	UnimplementedRpcServer
	clients          map[string]grpc.BidiStreamingServer[JoinRequest, JoinReply]
	multipassClients map[string]agent.Client

	clientsMu          sync.RWMutex
	multipassClientsMu sync.RWMutex
}

type Server interface {
	GetMultipassClients() map[string]agent.Client
}

func (s *server) GetMultipassClients() map[string]agent.Client {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()
	return s.multipassClients
}

func (s *server) addClient(uid string, stream grpc.BidiStreamingServer[JoinRequest, JoinReply]) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()
	s.clients[uid] = stream
}

func (s *server) removeClient(uid string) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()
	delete(s.clients, uid)
}

func (s *server) addMultipassClient(uid string, client agent.Client) {
	s.multipassClientsMu.Lock()
	defer s.multipassClientsMu.Unlock()
	s.multipassClients[uid] = client
}

func (s *server) removeMultipassClient(uid string) {
	s.multipassClientsMu.Lock()
	defer s.multipassClientsMu.Unlock()
	err := s.multipassClients[uid].Close()
	if err != nil {
		log.Fatalf("failed to close multipass client: %v", err)
	}
	delete(s.multipassClients, uid)
}

func (s *server) Join(stream grpc.BidiStreamingServer[JoinRequest, JoinReply]) error {
	id := uuid.Must(uuid.NewRandom()).String()
	defer log.Printf("client disconnected: %s", id)
	defer s.removeClient(id)
	defer s.removeMultipassClient(id)

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

		s.addClient(id, stream)
		s.addMultipassClient(id, agentClient)

		log.Printf("joined hostname: %s", request.Hostname)
	}
}

func NewServer(addr string, clusterServerCh chan Server) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := &server{
		clients:          make(map[string]grpc.BidiStreamingServer[JoinRequest, JoinReply]),
		clientsMu:        sync.RWMutex{},
		multipassClients: make(map[string]agent.Client),
	}
	clusterServerCh <- server
	RegisterRpcServer(grpcServer, server)
	return grpcServer.Serve(lis)
}
