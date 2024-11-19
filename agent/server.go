package agent

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"multipass-cluster/multipass"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	if listReplyInstanceList, ok := listContents.(*multipass.ListReply_InstanceList); !ok {
		instances := listReplyInstanceList.InstanceList.GetInstances()
		for _, instance := range instances {
			response.Names = append(response.Names, instance.Name)
		}
	}

	return response, nil
}

type shellReqReader struct {
	stream grpc.BidiStreamingServer[ShellRequest, ShellReply]
	done   bool
}

func (s *shellReqReader) Read(p []byte) (n int, err error) {
	if s.done {
		return 0, io.EOF
	}

	in, err := s.stream.Recv()
	if err == io.EOF {
		s.done = true
		return 0, io.EOF
	}

	if err != nil {
		return 0, err
	}

	n = copy(p, in.GetInBuffer())
	return n, nil
}

type shellReqWriter struct {
	stream grpc.BidiStreamingServer[ShellRequest, ShellReply]
}

func (s *shellReqWriter) Write(p []byte) (n int, err error) {
	err = s.stream.Send(&ShellReply{OutBuffer: p})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func (s *server) Shell(stream grpc.BidiStreamingServer[ShellRequest, ShellReply]) error {
	outW := &shellReqWriter{stream: stream}
	errW := &shellReqWriter{stream: stream}
	inR := &shellReqReader{stream: stream}

	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}

		err = s.ssh(stream.Context(), in.GetInstanceName(), int(in.GetHeight()), int(in.GetWidth()), outW, errW, inR)
		if err != nil {
			return err
		}
	}
}

//nolint:funlen
func (s *server) ssh(ctx context.Context, instanceName string, height int, width int, outW io.Writer, errW io.Writer, inR io.Reader) error {
	stream, err := s.multipassClient.SshInfo(ctx)
	if err != nil {
		return err
	}

	res, err := streamHandler[multipass.SSHInfoRequest, multipass.SSHInfoReply](
		stream,
		&multipass.SSHInfoRequest{
			InstanceName: []string{instanceName},
		},
	)
	if err != nil {
		return err
	}

	info, ok := res.SshInfo[instanceName]
	if !ok {
		return fmt.Errorf("instance not found: %s", instanceName)
	}

	signer, err := ssh.ParsePrivateKey([]byte(info.PrivKeyBase64))
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: info.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // nolint:gosec
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", info.Host, info.Port), config)
	if err != nil {
		return err
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(sshClient)

	sshSession, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
	}(sshSession)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	term := os.Getenv("TERM")
	if term == "" {
		term = "xterm"
	}

	if err := sshSession.RequestPty(term, height, width, modes); err != nil {
		return err
	}

	sshSession.Stdout = outW
	sshSession.Stderr = errW
	sshSession.Stdin = inR

	if err := sshSession.Shell(); err != nil {
		return err
	}

	if err := sshSession.Wait(); err != nil {
		var e *ssh.ExitError
		if errors.As(err, &e) && e.ExitStatus() == 130 {
			return nil
		}

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
