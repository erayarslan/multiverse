package agent

import (
	"context"

	"github.com/erayarslan/multiverse/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	conn   *grpc.ClientConn
	client RpcClient
}

type Client interface {
	Instances(ctx context.Context) (*GetInstancesReply, error)
	Info(ctx context.Context) (*GetInfoReply, error)
	Shell(ctx context.Context) (grpc.BidiStreamingClient[common.ShellRequest, common.ShellReply], error)
	Close() error
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Instances(ctx context.Context) (*GetInstancesReply, error) {
	return c.client.Instances(ctx, &GetInstancesRequest{})
}

func (c *client) Info(ctx context.Context) (*GetInfoReply, error) {
	return c.client.Info(ctx, &GetInfoRequest{})
}

func (c *client) Shell(ctx context.Context) (grpc.BidiStreamingClient[common.ShellRequest, common.ShellReply], error) {
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
