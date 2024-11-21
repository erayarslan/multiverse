package agent

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
	Shell(ctx context.Context) (grpc.BidiStreamingClient[ShellRequest, ShellReply], error)
	Close() error
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) List(ctx context.Context) ([]string, error) {
	names := make([]string, 0)

	response, err := c.client.List(ctx, &ListRequest{})
	if err != nil {
		return names, err
	}

	return response.GetNames(), err
}

func (c *client) Shell(ctx context.Context) (grpc.BidiStreamingClient[ShellRequest, ShellReply], error) {
	return c.client.Shell(ctx)
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
