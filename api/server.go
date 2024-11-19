package api

import (
	"context"
	"multipass-cluster/cluster"
	"net"

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
	var allNames []string
	for _, client := range *s.clusterServer.GetClients() {
		names, err := client.List(ctx)
		if err != nil {
			return nil, err
		}

		allNames = append(allNames, names...)
	}

	return &ListReply{
		Names: allNames,
	}, nil
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
