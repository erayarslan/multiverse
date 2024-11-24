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
	Instances(ctx context.Context, request *GetInstancesRequest) (*GetInstancesReply, error)
	Info(ctx context.Context, request *common.GetInfoRequest) (*common.GetInfoReply, error)
	Shell(ctx context.Context) (grpc.BidiStreamingClient[common.ShellRequest, common.ShellReply], error)
	Close() error
	Launch(ctx context.Context, request *common.LaunchRequest) (*common.LaunchReply, error)
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Instances(ctx context.Context, request *GetInstancesRequest) (*GetInstancesReply, error) {
	return c.client.Instances(ctx, request)
}

func (c *client) Info(ctx context.Context, request *common.GetInfoRequest) (*common.GetInfoReply, error) {
	return c.client.Info(ctx, request)
}

func (c *client) Launch(ctx context.Context, launchRequest *common.LaunchRequest) (*common.LaunchReply, error) {
	return c.client.Launch(ctx, launchRequest)
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
