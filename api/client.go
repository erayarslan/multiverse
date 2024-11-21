package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"multiverse/common"
	"os"
	"os/signal"
	"syscall"

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
	width  int
	height int
}

func (s *shellRequestWriter) Write(p []byte) (n int, err error) {
	err = s.stream.Send(&ShellRequest{InBuffer: p, Width: int64(s.width), Height: int64(s.height)})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func NewShellRequestWriter(
	stream grpc.BidiStreamingClient[ShellRequest, ShellReply],
	width int, height int, stdOutFd int,
) (io.Writer, error) {
	writer := &shellRequestWriter{width: width, height: height, stream: stream}

	go func(writer *shellRequestWriter) {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGWINCH)

		for range ch {
			if newWidth, newHeight, err := term.GetSize(stdOutFd); err == nil {
				writer.width = newWidth
				writer.height = newHeight
				_, _ = writer.Write([]byte{})
			}
		}
	}(writer)

	return writer, nil
}

func (c *client) Shell(ctx context.Context, nodeName string, instanceName string) error {
	stdInFd := int(os.Stdin.Fd())
	stdOutFd := int(os.Stdout.Fd())

	width, height, err := term.GetSize(stdOutFd)
	if err != nil {
		return nil
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(
		"nodeName", nodeName,
		"instanceName", instanceName,
		"width", fmt.Sprintf("%d", width),
		"height", fmt.Sprintf("%d", height),
	))

	stream, err := c.client.Shell(ctx)
	if err != nil {
		return err
	}

	state, err := term.MakeRaw(stdInFd)
	if err != nil {
		return err
	}
	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {
			panic(err)
		}
		log.Printf("restored terminal state")
	}(stdInFd, state)

	stdin, err := NewShellRequestWriter(stream, width, height, stdOutFd)
	if err != nil {
		return err
	}

	go func() {
		var err error
		if _, err = io.Copy(stdin, os.Stdin); err != nil {
			log.Printf("failed to copy stdin: %v", err)
		}
	}()

	return common.ListenBidiClient(stream, func(res *ShellReply) error {
		var err error
		if _, err = os.Stdout.Write(res.GetOutBuffer()); err != nil {
			return err
		}
		if _, err = os.Stderr.Write(res.GetErrBuffer()); err != nil {
			return err
		}
		return nil
	})
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
