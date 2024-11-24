package cluster

import (
	"fmt"
	"log"
	"multiverse/agent"
	"multiverse/common"
	"net"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type WorkerInfo struct {
	AgentClient agent.Client
	IPPort      string
	Stream      grpc.BidiStreamingServer[SyncRequest, SyncReply]
	State       *State
	LastSync    *timestamppb.Timestamp
	NodeName    string
	UUID        string
}

type server struct {
	UnimplementedRpcServer
	listener      net.Listener
	workerInfoMap map[string]*WorkerInfo
	grpcServer    *grpc.Server
	workersMu     sync.RWMutex
}

type Server interface {
	GetWorkerInfoMap() map[string]*WorkerInfo
	Serve() error
}

func (s *server) GetWorkerInfoMap() map[string]*WorkerInfo {
	s.workersMu.Lock()
	defer s.workersMu.Unlock()
	return s.workerInfoMap
}

func (s *server) addWorkerInfo(uid string, workerInfo *WorkerInfo) {
	s.workersMu.Lock()
	defer s.workersMu.Unlock()
	s.workerInfoMap[uid] = workerInfo
}

func (s *server) updateState(uid string, state *State) error {
	s.workersMu.Lock()
	defer s.workersMu.Unlock()
	if workerInfo, ok := s.workerInfoMap[uid]; ok {
		workerInfo.State = state
		workerInfo.LastSync = timestamppb.Now()
	}
	return nil
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

func (s *server) Sync(stream grpc.BidiStreamingServer[SyncRequest, SyncReply]) error {
	id := uuid.Must(uuid.NewRandom()).String()
	defer log.Printf("client disconnected: %s", id)
	defer s.removeWorkerInfo(id)

	ctx := stream.Context()

	p, ok := peer.FromContext(ctx)
	if !ok {
		return fmt.Errorf("failed to get peer from context")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("metadata not found in context")
	}

	nodeName := md.Get("nodeName")
	if len(nodeName) == 0 {
		return fmt.Errorf("node name not found in context")
	}

	agentPort := md.Get("agentPort")
	if len(agentPort) == 0 {
		return fmt.Errorf("agent port not found in context")
	}
	port, _ := strconv.Atoi(agentPort[0])

	host := strings.Split(p.Addr.String(), ":")[0]
	target := fmt.Sprintf("%s:%d", host, port)

	if err := stream.Send(&SyncReply{
		Uuid: id,
	}); err != nil {
		return fmt.Errorf("failed to send join reply on master: %w", err)
	}

	agentClient, err := agent.NewClient(target)
	if err != nil {
		return fmt.Errorf("failed to create multipass client: %w", err)
	}

	s.addWorkerInfo(id, &WorkerInfo{
		AgentClient: agentClient,
		Stream:      stream,
		NodeName:    nodeName[0],
		UUID:        id,
		State: &State{
			Instances: make([]*Instance, 0),
		},
		IPPort:   target,
		LastSync: timestamppb.Now(),
	})

	log.Printf("joined node name: %s, uuid: %s", nodeName[0], id)

	return common.ListenBidiServer(stream, func(req *SyncRequest) error {
		return s.updateState(id, req.GetState())
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
		workerInfoMap: map[string]*WorkerInfo{},
		listener:      lis,
		grpcServer:    grpcServer,
	}
	RegisterRpcServer(grpcServer, server)
	return server, nil
}
