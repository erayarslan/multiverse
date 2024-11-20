package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"multiverse/agent"
	"multiverse/cluster"
	"net"

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
	for _, worker := range *s.clusterServer.GetWorkers() {
		names, err := worker.AgentClient.List(ctx)
		if err != nil {
			return nil, err
		}

		for _, name := range names {
			listReply.Instances = append(listReply.Instances, &Instance{
				NodeName:     worker.NodeName,
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
	for _, worker := range *s.clusterServer.GetWorkers() {
		if worker.NodeName == nodeName[0] {
			agentClient = worker.AgentClient
			break
		}
	}
	if agentClient == nil {
		return fmt.Errorf("agent client not found on node name: %s", nodeName[0])
	}

	agentStream, err := agentClient.Shell(stream.Context())
	if err != nil {
		return err
	}

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				return
			}

			if err = agentStream.Send(&agent.ShellRequest{
				InBuffer: in.GetInBuffer(),
			}); err != nil {
				log.Printf("failed to send shell request: %v", err)
				return
			}
		}
	}()

	for {
		in, err := agentStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if err = stream.Send(&ShellReply{
			OutBuffer: in.GetOutBuffer(),
			ErrBuffer: in.GetErrBuffer(),
		}); err != nil {
			return err
		}
	}

	return nil
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
