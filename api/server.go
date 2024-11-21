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

func (s *server) List(ctx context.Context, _ *ListRequest) (*ListReply, error) {
	listReply := &ListReply{
		Instances: make([]*Instance, 0),
	}
	for _, workerInfo := range s.clusterServer.GetWorkerInfoMap() {
		names, err := workerInfo.AgentClient.List(ctx)
		if err != nil {
			return nil, err
		}

		for _, name := range names {
			listReply.Instances = append(listReply.Instances, &Instance{
				NodeName:     workerInfo.NodeName,
				InstanceName: name,
			})
		}
	}

	return listReply, nil
}

func (s *server) Shell(stream grpc.BidiStreamingServer[ShellRequest, ShellReply]) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("metadata not found in context")
	}

	nodeName := md.Get("nodeName")
	if len(nodeName) == 0 {
		return fmt.Errorf("node name not found in context")
	}

	instanceName := md.Get("instanceName")
	if len(instanceName) == 0 {
		return fmt.Errorf("instance name not found in context")
	}

	var agentClient agent.Client
	for _, workerInfo := range s.clusterServer.GetWorkerInfoMap() {
		if workerInfo.NodeName == nodeName[0] {
			agentClient = workerInfo.AgentClient
			break
		}
	}
	if agentClient == nil {
		return fmt.Errorf("agent client not found on node name: %s", nodeName[0])
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
