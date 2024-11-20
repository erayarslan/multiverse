package agent

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"multipass-cluster/multipass"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type Info struct {
	Port int
}

type Server interface {
	Serve() error
	Port() int
}

type server struct {
	UnimplementedRpcServer
	multipassClient multipass.RpcClient
	listener        net.Listener
	grpcServer      *grpc.Server
}

func streamHandler[Req any, Res any](stream grpc.BidiStreamingClient[Req, Res], req *Req) (res *Res, err error) {
	err = stream.Send(req)
	if err != nil {
		return nil, err
	}

	in, err := stream.Recv()
	if err == io.EOF {
		err = stream.CloseSend()
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) Port() int {
	return s.listener.Addr().(*net.TCPAddr).Port
}

func (s *server) List(ctx context.Context, _ *ListRequest) (*ListReply, error) {
	stream, err := s.multipassClient.List(ctx)
	if err != nil {
		return nil, err
	}

	res, err := streamHandler[multipass.ListRequest, multipass.ListReply](stream, &multipass.ListRequest{})
	if err != nil {
		return nil, err
	}

	response := &ListReply{
		Names: make([]string, 0),
	}

	listContents := res.GetListContents()
	if listContents == nil {
		return response, nil
	}

	if listReplyInstanceList, ok := listContents.(*multipass.ListReply_InstanceList); ok {
		instances := listReplyInstanceList.InstanceList.GetInstances()
		for _, instance := range instances {
			response.Names = append(response.Names, instance.Name)
		}
	}

	return response, nil
}

type shellRequestReader struct {
	stream grpc.BidiStreamingServer[ShellRequest, ShellReply]
}

func (s *shellRequestReader) Read(p []byte) (n int, err error) {
	in, err := s.stream.Recv()
	if err != nil {
		return 0, err
	}

	n = copy(p, in.GetInBuffer())
	return n, nil
}

type shellReplyWriter struct {
	stream grpc.BidiStreamingServer[ShellRequest, ShellReply]
	isErr  bool
}

func (s *shellReplyWriter) Write(p []byte) (n int, err error) {
	reply := &ShellReply{}
	if s.isErr {
		reply.ErrBuffer = p
	} else {
		reply.OutBuffer = p
	}
	err = s.stream.Send(reply)
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func (s *server) Shell(stream grpc.BidiStreamingServer[ShellRequest, ShellReply]) error {
	stdout := &shellReplyWriter{stream: stream, isErr: false}
	stderr := &shellReplyWriter{stream: stream, isErr: true}
	stdin := &shellRequestReader{stream: stream}

	sshInfoStream, err := s.multipassClient.SshInfo(context.Background())
	if err != nil {
		return err
	}

	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("metadata not found in context")
	}

	instanceName := md.Get("instanceName")
	if len(instanceName) == 0 {
		return fmt.Errorf("instance name not found in context")
	}

	res, err := streamHandler[multipass.SSHInfoRequest, multipass.SSHInfoReply](
		sshInfoStream,
		&multipass.SSHInfoRequest{
			InstanceName: []string{instanceName[0]},
		},
	)
	if err != nil {
		return err
	}

	height := md.Get("height")
	if len(height) == 0 {
		return fmt.Errorf("height not found in context")
	}

	width := md.Get("width")
	if len(width) == 0 {
		return fmt.Errorf("width not found in context")
	}

	info, ok := res.SshInfo[instanceName[0]]
	if !ok {
		return fmt.Errorf("instance not found: %s", instanceName)
	}

	h, _ := strconv.Atoi(height[0])
	w, _ := strconv.Atoi(width[0])

	ssh := NewSSH(info.Host, int(info.Port), info.Username, []byte(info.PrivKeyBase64), stdout, stderr, stdin, h, w)
	if err = ssh.Start(); err != nil {
		return err
	}

	return nil
}

func NewServer(target string, addr string, multipassCertFilePath string, multipassKeyFilePath string) (Server, error) {
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

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:0", addr))
	if err != nil {
		return nil, err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := &server{
		multipassClient: multipass.NewRpcClient(conn),
		listener:        lis,
		grpcServer:      grpcServer,
	}
	RegisterRpcServer(grpcServer, server)
	return server, nil
}
