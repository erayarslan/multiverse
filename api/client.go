package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type client struct {
	conn   *grpc.ClientConn
	client RpcClient
}

type Client interface {
	List(ctx context.Context) (*ListReply, error)
	Shell(ctx context.Context, nodeName string, instanceName string) error
	Close() error
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) List(ctx context.Context) (*ListReply, error) {
	return c.client.List(ctx, &ListRequest{})
}

type shellRequestWriter struct {
	stream grpc.BidiStreamingClient[ShellRequest, ShellReply]
}

func (s *shellRequestWriter) Write(p []byte) (n int, err error) {
	err = s.stream.Send(&ShellRequest{InBuffer: p})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func (c *client) Shell(ctx context.Context, nodeName string, instanceName string) error {
	fdIn := int(os.Stdin.Fd())
	fdOut := int(os.Stdout.Fd())

	w, h, err := term.GetSize(fdOut)
	if err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(
		"nodeName", nodeName,
		"instanceName", instanceName,
		"width", fmt.Sprintf("%d", w),
		"height", fmt.Sprintf("%d", h),
	))

	stream, err := c.client.Shell(ctx)
	if err != nil {
		return err
	}

	state, err := term.MakeRaw(fdIn)
	if err != nil {
		return err
	}
	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {
			panic(err)
		}
		log.Printf("restored terminal state")
	}(fdIn, state)

	stdin := &shellRequestWriter{stream: stream}

	go func() {
		_, err := io.Copy(stdin, os.Stdin)
		if err != nil {
			log.Printf("failed to copy stdin: %v", err)
		}
	}()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(in.GetOutBuffer())
		if err != nil {
			return err
		}

		_, err = os.Stderr.Write(in.GetErrBuffer())
		if err != nil {
			return err
		}
	}

	return nil
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
