package cluster

import (
	"context"
	"fmt"
	"log"
	"multiverse/agent"
	"multiverse/common"
	"multiverse/multipass"
	"strconv"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	stream          grpc.BidiStreamingClient[SyncRequest, SyncReply]
	client          RpcClient
	agentServer     agent.Server
	multipassClient multipass.Client
	state           agent.State
	conn            *grpc.ClientConn
	uuid            string
	nodeName        string
	closed          bool
}

type Client interface {
	Sync() error
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

func (c *client) Sync() error {
	if !c.isReady() {
		return fmt.Errorf("could not connect")
	}
	if err := c.sync(); err != nil && !c.closed {
		log.Printf("error while sync: %v", err)
		log.Printf("reconnecting...")
		return c.Sync()
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

func (c *client) sync() error {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"nodeName", c.nodeName,
		"agentPort", strconv.Itoa(c.agentServer.Port()),
	))

	var err error
	c.stream, err = c.client.Sync(ctx)
	if err != nil {
		return err
	}

	return common.ListenBidiClient(c.stream, func(res *SyncReply) error {
		c.uuid = res.Uuid
		log.Printf("joined with uuid: %s", c.uuid)
		return nil
	})
}

func (c *client) stateSync() {
	for state := range c.state.Listen() { //nolint:govet
		if c.closed {
			continue
		}

		clusterInstances := make([]*Instance, 0, len(state.Instances))
		for _, instance := range state.Instances {
			clusterInstances = append(clusterInstances, &Instance{
				Name:  instance.Name,
				State: instance.State,
				Ipv4:  instance.Ipv4,
				Image: instance.Image,
			})
		}

		if c.stream == nil {
			continue
		}

		if err := c.stream.Send(&SyncRequest{
			State: &State{
				Instances: clusterInstances,
			},
		}); err != nil {
			log.Printf("error while sending state: %v", err)
		}
	}
}

func NewClient(addr string, nodeName string,
	agentServer agent.Server, multipassClient multipass.Client, state agent.State,
) (Client, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}
	rpcClient := NewRpcClient(conn)
	c := &client{
		conn:            conn,
		client:          rpcClient,
		agentServer:     agentServer,
		nodeName:        nodeName,
		multipassClient: multipassClient,
		state:           state,
	}
	go c.stateSync()
	return c, nil
}
