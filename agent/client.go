package agent

import (
	"context"
	"golang.org/x/crypto/ssh/terminal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
)

type client struct {
	conn   *grpc.ClientConn
	client RpcClient
}

type Client interface {
	List(ctx context.Context) ([]string, error)
	Shell(ctx context.Context, instanceName string) error
	Close() error
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) List(ctx context.Context) ([]string, error) {
	response, err := c.client.List(ctx, &ListRequest{})

	return response.GetNames(), err
}

func (c *client) Shell(ctx context.Context, instanceName string) error {
	fdOut := int(os.Stdout.Fd())
	_ = int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fdOut)
	if err != nil {
		return err
	}
	defer func(fd int, oldState *terminal.State) {
		err := terminal.Restore(fd, oldState)
		if err != nil {
			panic(err)
		}
	}(fdOut, state)

	w, h, err := terminal.GetSize(fdOut)
	if err != nil {
		return err
	}

	stream, err := c.client.Shell(ctx)
	if err != nil {
		return err
	}

	err = stream.Send(&ShellRequest{
		InstanceName: instanceName,
		Height:       uint32(h),
		Width:        uint32(w),
		InBuffer:     make([]byte, 0),
	})

	if err != nil {
		return err
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("received on worker: %s", in)
	}

	err = stream.CloseSend()

	return nil
}

func NewClient(addr string) (Client, error) {
	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
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
