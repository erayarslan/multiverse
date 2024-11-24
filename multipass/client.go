package multipass

import (
	"context"
	"crypto/tls"
	"fmt"
	"multiverse/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type client struct {
	rpcClient RpcClient
}

type instance struct {
	Name  string
	State string
	Image string
	Ipv4  []string
}

type Client interface {
	List(ctx context.Context) ([]*instance, error)
	SSHInfo(ctx context.Context, instanceName string) (*SSHInfo, error)
}

func (s InstanceStatus_Status) ToString() string {
	switch s {
	case InstanceStatus_RUNNING:
		return "Running"
	case InstanceStatus_STOPPED:
		return "Stopped"
	case InstanceStatus_DELETED:
		return "Deleted"
	case InstanceStatus_STARTING:
		return "Starting"
	case InstanceStatus_RESTARTING:
		return "Restarting"
	case InstanceStatus_DELAYED_SHUTDOWN:
		return "Delayed Shutdown"
	case InstanceStatus_SUSPENDING:
		return "Suspending"
	case InstanceStatus_SUSPENDED:
		return "Suspended"
	case InstanceStatus_UNKNOWN:
		return "Unknown"
	default:
		return "Unknown"
	}
}

func (c *client) List(ctx context.Context) ([]*instance, error) {
	stream, err := c.rpcClient.List(ctx)
	if err != nil {
		return nil, err
	}

	res, err := common.ExecuteOnceWithBidiClient(stream, &ListRequest{
		VerbosityLevel: 1,
		RequestIpv4:    true,
	})
	if err != nil {
		return nil, err
	}

	instances := make([]*instance, 0)

	listContents := res.GetListContents()
	if listContents == nil {
		return instances, nil
	}

	if listReplyInstanceList, ok := listContents.(*ListReply_InstanceList); ok {
		multipassInstances := listReplyInstanceList.InstanceList.GetInstances()
		for _, multipassInstance := range multipassInstances {
			instances = append(instances, &instance{
				Name:  multipassInstance.Name,
				State: multipassInstance.InstanceStatus.Status.ToString(),
				Ipv4:  multipassInstance.Ipv4,
				Image: fmt.Sprintf("Ubuntu %s", multipassInstance.CurrentRelease),
			})
		}
	}

	return instances, nil
}

func (c *client) SSHInfo(ctx context.Context, instanceName string) (*SSHInfo, error) {
	stream, err := c.rpcClient.SshInfo(ctx)
	if err != nil {
		return nil, err
	}

	res, err := common.ExecuteOnceWithBidiClient(stream, &SSHInfoRequest{InstanceName: []string{instanceName}})
	if err != nil {
		return nil, err
	}

	info, ok := res.SshInfo[instanceName]
	if !ok {
		return nil, fmt.Errorf("instance not found: %s", instanceName)
	}

	return info, nil
}

func NewClient(target string, multipassCertFilePath string, multipassKeyFilePath string) (Client, error) {
	multipassCertificate, err := tls.LoadX509KeyPair(multipassCertFilePath, multipassKeyFilePath)
	if err != nil {
		return nil, err
	}

	multipassTransportCredentials := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{multipassCertificate},
		InsecureSkipVerify: true, // nolint:gosec
	})

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(multipassTransportCredentials)}
	conn, err := grpc.NewClient(target, dialOpts...)
	if err != nil {
		return nil, err
	}

	return &client{
		rpcClient: NewRpcClient(conn),
	}, nil
}
