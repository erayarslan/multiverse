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
	List(ctx context.Context) ([]*Instance, error)
	Shell(ctx context.Context) (grpc.BidiStreamingClient[ShellRequest, ShellReply], error)
	Close() error
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) List(ctx context.Context) ([]*Instance, error) {
	response, err := c.client.List(ctx, &ListRequest{})
	if err != nil {
		return make([]*Instance, 0), err
	}

	instances := make([]*Instance, len(response.GetInstances()))
	for i, agentInstance := range response.GetInstances() {
		instances[i] = &Instance{
			Name:  agentInstance.Name,
			State: agentInstance.State,
			Ipv4:  agentInstance.Ipv4,
			Image: agentInstance.Image,
		}
	}

	return instances, nil
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
