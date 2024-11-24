package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/erayarslan/multiverse/agent"
	"github.com/erayarslan/multiverse/cluster"
	"github.com/erayarslan/multiverse/common"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

type server struct {
	UnimplementedRpcServer
	clusterServer cluster.Server
	listener      net.Listener
	grpcServer    *grpc.Server
}

type Server interface {
	Serve() error
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) Info(ctx context.Context, _ *GetInfoRequest) (*GetInfoReply, error) {
	getInfoReply := &GetInfoReply{
		Instances: make([]*GetInfoInstance, 0),
	}

	s.clusterServer.IterateWorkers(func(workerInfo *cluster.WorkerInfo) bool {
		info, err := workerInfo.AgentClient.Info(ctx, &common.GetInfoRequest{})
		if err != nil {
			log.Printf("failed to get info: %v", err)
			return true
		}

		for _, instance := range info.Instances {
			getInfoReply.Instances = append(getInfoReply.Instances, &GetInfoInstance{
				NodeName: workerInfo.NodeName,
				Instance: instance,
			})
		}

		return true
	})

	return getInfoReply, nil
}

func (s *server) Instances(_ context.Context, _ *GetInstancesRequest) (*GetInstancesReply, error) {
	getInstancesReply := &GetInstancesReply{
		Instances: make([]*Instance, 0),
	}

	s.clusterServer.IterateWorkers(func(workerInfo *cluster.WorkerInfo) bool {
		for _, instance := range workerInfo.State.Instances {
			getInstancesReply.Instances = append(getInstancesReply.Instances, &Instance{
				NodeName: workerInfo.NodeName,
				Instance: instance,
			})
		}
		return true
	})

	return getInstancesReply, nil
}

func (s *server) Nodes(_ context.Context, _ *GetNodesRequest) (*GetNodesReply, error) {
	getNodesReply := &GetNodesReply{
		Nodes: make([]*Node, 0),
	}

	s.clusterServer.IterateWorkers(func(workerInfo *cluster.WorkerInfo) bool {
		getNodesReply.Nodes = append(getNodesReply.Nodes, &Node{
			Name:     workerInfo.NodeName,
			LastSync: workerInfo.LastSync,
			Ipv4:     []string{workerInfo.IPPort},
			Resource: workerInfo.State.Resource,
		})
		return true
	})

	return getNodesReply, nil
}

func (s *server) Launch(ctx context.Context, req *common.LaunchRequest) (*common.LaunchReply, error) {
	// todo: detect best worker due to resource
	var agentClient agent.Client
	s.clusterServer.IterateWorkers(func(workerInfo *cluster.WorkerInfo) bool {
		agentClient = workerInfo.AgentClient
		return false
	})
	if agentClient == nil {
		return nil, fmt.Errorf("agent client not found")
	}
	return agentClient.Launch(ctx, req)
}

func (s *server) Shell(stream grpc.BidiStreamingServer[common.ShellRequest, common.ShellReply]) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("metadata not found in context")
	}

	instanceName := md.Get("instanceName")
	if len(instanceName) == 0 {
		return fmt.Errorf("instance name not found in context")
	}

	var agentClient agent.Client
	s.clusterServer.IterateWorkers(func(workerInfo *cluster.WorkerInfo) bool {
		for _, instance := range workerInfo.State.Instances {
			if instance.Name == instanceName[0] {
				agentClient = workerInfo.AgentClient
				return false
			}
		}
		return true
	})
	if agentClient == nil {
		return fmt.Errorf("agent client not found with instance name: %s", instanceName[0])
	}

	ctx := metadata.NewOutgoingContext(context.Background(), md.Copy())
	agentStream, err := agentClient.Shell(ctx)
	if err != nil {
		return err
	}

	go func() {
		err := common.ListenBidiServer(stream, func(req *common.ShellRequest) error {
			return agentStream.Send(&common.ShellRequest{
				InBuffer: req.GetInBuffer(),
				Width:    req.GetWidth(),
				Height:   req.GetHeight(),
			})
		})
		if err != nil {
			if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
				return
			} else {
				log.Printf("failed to listen stream: %v", err)
			}
		}
	}()

	return common.ListenBidiClient(agentStream, func(res *common.ShellReply) error {
		return stream.Send(&common.ShellReply{
			OutBuffer: res.GetOutBuffer(),
			ErrBuffer: res.GetErrBuffer(),
		})
	})
}

func NewServer(addr string, clusterServer cluster.Server) (Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := &server{
		clusterServer: clusterServer,
		listener:      lis,
		grpcServer:    grpcServer,
	}
	RegisterRpcServer(grpcServer, server)
	return server, nil
}
