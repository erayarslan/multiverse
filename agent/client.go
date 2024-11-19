package agent

import (
	"context"
	"io"
	"log"
	"os"

	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	state, err := term.MakeRaw(fdOut)
	if err != nil {
		return err
	}
	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {
			panic(err)
		}
	}(fdOut, state)

	w, h, err := term.GetSize(fdOut)
	if err != nil {
		return err
	}

	stream, err := c.client.Shell(ctx)
	if err != nil {
		return err
	}

	err = stream.Send(&ShellRequest{
		InstanceName: instanceName,
		Height:       int64(h),
		Width:        int64(w),
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

	return stream.CloseSend()
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
