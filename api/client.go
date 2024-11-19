package api

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	conn   *grpc.ClientConn
	client RpcClient
}

type Client interface {
	List(ctx context.Context) ([]string, error)
	Close() error
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) List(ctx context.Context) ([]string, error) {
	response, err := c.client.List(ctx, &ListRequest{})

	return response.GetNames(), err
}

func NewClient(addr string) (Client, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}
	rpcClient := NewRpcClient(conn)
	return &client{
		conn:   conn,
		client: rpcClient,
	}, nil
}
