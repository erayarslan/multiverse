package api

import (
	"context"
	"fmt"
	"log"
	"multiverse/agent"
	"multiverse/cluster"
	"multiverse/common"
	"net"

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

func (s *server) Instances(_ context.Context, _ *GetInstancesRequest) (*GetInstancesReply, error) {
	getInstancesReply := &GetInstancesReply{
		Instances: make([]*Instance, 0),
	}
	for _, workerInfo := range s.clusterServer.GetWorkerInfoMap() {
		for _, instance := range workerInfo.State.Instances {
			getInstancesReply.Instances = append(getInstancesReply.Instances, &Instance{
				NodeName:     workerInfo.NodeName,
				InstanceName: instance.Name,
				State:        instance.State,
				Ipv4:         instance.Ipv4,
				Image:        instance.Image,
			})
		}
	}

	return getInstancesReply, nil
}

func (s *server) Nodes(_ context.Context, _ *GetNodesRequest) (*GetNodesReply, error) {
	getNodesReply := &GetNodesReply{
		Nodes: make([]*Node, 0),
	}
	for _, workerInfo := range s.clusterServer.GetWorkerInfoMap() {
		getNodesReply.Nodes = append(getNodesReply.Nodes, &Node{
			Name:     workerInfo.NodeName,
			LastSync: workerInfo.LastSync,
			Ipv4:     []string{workerInfo.IPPort},
		})
	}

	return getNodesReply, nil
}

func (s *server) Shell(stream grpc.BidiStreamingServer[ShellRequest, ShellReply]) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("metadata not found in context")
	}

	instanceName := md.Get("instanceName")
	if len(instanceName) == 0 {
		return fmt.Errorf("instance name not found in context")
	}

	var agentClient agent.Client
	for _, workerInfo := range s.clusterServer.GetWorkerInfoMap() {
		for _, instance := range workerInfo.State.Instances {
			if instance.Name == instanceName[0] {
				agentClient = workerInfo.AgentClient
				break
			}
		}
	}
	if agentClient == nil {
		return fmt.Errorf("agent client not found with instance name: %s", instanceName[0])
	}

	ctx := metadata.NewOutgoingContext(context.Background(), md.Copy())
	agentStream, err := agentClient.Shell(ctx)
	if err != nil {
		return err
	}

	go func() {
		err := common.ListenBidiServer(stream, func(req *ShellRequest) error {
			return agentStream.Send(&agent.ShellRequest{
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

	return common.ListenBidiClient(agentStream, func(res *agent.ShellReply) error {
		return stream.Send(&ShellReply{
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
