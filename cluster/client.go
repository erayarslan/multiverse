package cluster

import (
	"context"
	"fmt"
	"log"
	"multipass-cluster/agent"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	conn   *grpc.ClientConn
	client RpcClient
	port   int64
}

type Client interface {
	Join() error
	Close() error
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Join() error {
	if !c.isReady() {
		return fmt.Errorf("could not connect")
	}
	if err := c.join(); err != nil {
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
	stream, err := c.client.Join(context.Background())
	if err != nil {
		return err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	err = stream.Send(&JoinRequest{
		Hostname: hostname,
		AgentInfo: &AgentInfo{
			Port: c.port,
		},
	})
	if err != nil {
		return err
	}

	ctx := stream.Context()
	errCh := make(chan error, 1)

	go func() {
		for {
			response, err := stream.Recv()
			if err != nil {
				errCh <- err
				return
			} else {
				log.Printf("received on worker: %s", response)
			}
		}
	}()

	go func() {
		<-ctx.Done()
		errCh <- ctx.Err()
	}()

	return <-errCh
}

func NewClient(addr string, agentServer agent.Server) (Client, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}
	rpcClient := NewRpcClient(conn)
	return &client{
		conn:   conn,
		client: rpcClient,
		port:   int64(agentServer.Port()),
	}, nil
}
