package api

import (
	"context"
	"google.golang.org/grpc"
	"multipass-cluster/cluster"
	"net"
)

type server struct {
	UnimplementedRpcServer
	clusterServer cluster.Server
}

type Server interface {
}

func (s *server) List(ctx context.Context, _ *ListRequest) (*ListReply, error) {
	var allNames []string
	for _, client := range s.clusterServer.GetMultipassClients() {
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

func NewServer(addr string, clusterServerCh chan cluster.Server) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterRpcServer(grpcServer, &server{
		clusterServer: <-clusterServerCh,
	})
	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
