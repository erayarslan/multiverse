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

type Worker struct {
	AgentClient agent.Client
	Stream      grpc.BidiStreamingServer[JoinRequest, JoinReply]
	NodeName    string
	UUID        string
}

type server struct {
	UnimplementedRpcServer
	listener   net.Listener
	workers    *map[string]Worker
	grpcServer *grpc.Server
	workersMu  sync.RWMutex
}

type Server interface {
	GetWorkers() *map[string]Worker
	Serve() error
}

func (s *server) GetWorkers() *map[string]Worker {
	s.workersMu.Lock()
	defer s.workersMu.Unlock()
	return s.workers
}

func (s *server) addWorker(uid string, worker Worker) {
	s.workersMu.Lock()
	defer s.workersMu.Unlock()
	(*s.workers)[uid] = worker
}

func (s *server) removeWorker(uid string) {
	s.workersMu.Lock()
	defer s.workersMu.Unlock()
	err := (*s.workers)[uid].AgentClient.Close()
	if err != nil {
		log.Printf("failed to close agent client of worker: %v", err)
	}
	delete(*s.workers, uid)
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) Join(stream grpc.BidiStreamingServer[JoinRequest, JoinReply]) error {
	id := uuid.Must(uuid.NewRandom()).String()
	defer log.Printf("client disconnected: %s", id)
	defer s.removeWorker(id)

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

		s.addWorker(id, Worker{
			AgentClient: agentClient,
			Stream:      stream,
			NodeName:    request.NodeName,
			UUID:        id,
		})

		log.Printf("joined node name: %s", request.NodeName)
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
		workersMu:  sync.RWMutex{},
		workers:    &map[string]Worker{},
		listener:   lis,
		grpcServer: grpcServer,
	}
	RegisterRpcServer(grpcServer, server)
	return server, nil
}
