package cluster

import (
	"fmt"
	"log"
	"multiverse/agent"
	"multiverse/common"
	"net"
	"strings"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type WorkerInfo struct {
	AgentClient agent.Client
	Stream      grpc.BidiStreamingServer[JoinRequest, JoinReply]
	NodeName    string
	UUID        string
}

type server struct {
	UnimplementedRpcServer
	listener      net.Listener
	workerInfoMap map[string]WorkerInfo
	grpcServer    *grpc.Server
	workersMu     sync.RWMutex
}

type Server interface {
	GetWorkerInfoMap() map[string]WorkerInfo
	Serve() error
}

func (s *server) GetWorkerInfoMap() map[string]WorkerInfo {
	s.workersMu.Lock()
	defer s.workersMu.Unlock()
	return s.workerInfoMap
}

func (s *server) addWorkerInfo(uid string, workerInfo WorkerInfo) {
	s.workersMu.Lock()
	defer s.workersMu.Unlock()
	s.workerInfoMap[uid] = workerInfo
}

func (s *server) removeWorkerInfo(uid string) {
	s.workersMu.Lock()
	defer s.workersMu.Unlock()
	if workerInfo, ok := s.workerInfoMap[uid]; ok {
		err := workerInfo.AgentClient.Close()
		if err != nil {
			log.Printf("failed to close agent client of worker: %v", err)
		}
	}
	delete(s.workerInfoMap, uid)
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) Join(stream grpc.BidiStreamingServer[JoinRequest, JoinReply]) error {
	id := uuid.Must(uuid.NewRandom()).String()
	defer log.Printf("client disconnected: %s", id)
	defer s.removeWorkerInfo(id)

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("failed to get peer from context")
	}

	return common.ListenBidiServer(stream, func(req *JoinRequest) error {
		if err := stream.Send(&JoinReply{
			Uuid: id,
		}); err != nil {
			return fmt.Errorf("failed to send join reply on master: %w", err)
		}

		host := strings.Split(p.Addr.String(), ":")[0]
		target := fmt.Sprintf("%s:%d", host, req.AgentInfo.Port)
		agentClient, err := agent.NewClient(target)
		if err != nil {
			return fmt.Errorf("failed to create multipass client: %w", err)
		}

		s.addWorkerInfo(id, WorkerInfo{
			AgentClient: agentClient,
			Stream:      stream,
			NodeName:    req.NodeName,
			UUID:        id,
		})
		log.Printf("joined node name: %s", req.NodeName)
		return nil
	})
}

func NewServer(addr string) (Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := &server{
		workersMu:     sync.RWMutex{},
		workerInfoMap: map[string]WorkerInfo{},
		listener:      lis,
		grpcServer:    grpcServer,
	}
	RegisterRpcServer(grpcServer, server)
	return server, nil
}
