package cluster

import (
	"context"
	"fmt"
	"io"
	"log"
	"multiverse/agent"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	stream   grpc.BidiStreamingClient[JoinRequest, JoinReply]
	client   RpcClient
	conn     *grpc.ClientConn
	uuid     string
	nodeName string
	port     int64
	closed   bool
}

type Client interface {
	Join() error
	Close() error
}

func (c *client) Close() error {
	c.closed = true
	if c.stream != nil {
		if err := c.stream.CloseSend(); err != nil {
			log.Printf("error while closing stream: %v", err)
		}
	}
	return c.conn.Close()
}

func (c *client) Join() error {
	if !c.isReady() {
		return fmt.Errorf("could not connect")
	}
	if err := c.join(); err != nil && !c.closed {
		log.Printf("error while joining: %v", err)
		log.Printf("reconnecting...")
		return c.Join()
	}
	return nil
}

func (c *client) isReady() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			c.conn.Connect()
			if c.conn.GetState() == connectivity.Ready {
				return true
			}
		case <-ctx.Done():
			return false
		}
	}
}

func (c *client) join() error {
	var err error
	c.stream, err = c.client.Join(context.Background())
	if err != nil {
		return err
	}

	err = c.stream.Send(&JoinRequest{
		NodeName: c.nodeName,
		AgentInfo: &AgentInfo{
			Port: c.port,
		},
	})
	if err != nil {
		return err
	}

	for {
		response, err := c.stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		c.uuid = response.Uuid
	}

	return nil
}

func NewClient(addr string, nodeName string, agentServer agent.Server) (Client, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}
	rpcClient := NewRpcClient(conn)
	return &client{
		conn:     conn,
		client:   rpcClient,
		port:     int64(agentServer.Port()),
		nodeName: nodeName,
	}, nil
}
