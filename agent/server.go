package agent

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"multipass-cluster/multipass"
	"net"
	"os"
)

type Info struct {
	Port int
}

type Server interface {
}

type server struct {
	UnimplementedRpcServer
	multipassClient multipass.RpcClient
}

func streamHandler[Req any, Res any](stream grpc.BidiStreamingClient[Req, Res], req *Req) (res *Res, err error) {
	err = stream.Send(req)
	if err != nil {
		return nil, err
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		return in, nil
	}

	err = stream.CloseSend()
	return nil, err
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

	var response = &ListReply{
		Names: make([]string, 0),
	}

	listContents := res.GetListContents()
	if listContents == nil {
		return response, nil
	}

	switch listContents.(type) {
	case *multipass.ListReply_InstanceList:
		instances := listContents.(*multipass.ListReply_InstanceList).InstanceList.GetInstances()
		for _, instance := range instances {
			response.Names = append(response.Names, instance.Name)
		}
	case *multipass.ListReply_SnapshotList:
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
	var outW = &shellReqWriter{stream: stream}
	var errW = &shellReqWriter{stream: stream}
	var inR = &shellReqReader{stream: stream}

	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}

		err = s.ssh(stream.Context(), in.GetInstanceName(), int(in.GetHeight()), int(in.GetWidth()), outW, errW, inR)
		if err != nil {
			return err
		}

		return nil
	}
}

func (s *server) ssh(ctx context.Context, instanceName string, height int, width int, outW io.Writer, errW io.Writer, inR io.Reader) error {
	stream, err := s.multipassClient.SshInfo(ctx)
	if err != nil {
		return err
	}

	res, err := streamHandler[multipass.SSHInfoRequest, multipass.SSHInfoReply](stream, &multipass.SSHInfoRequest{InstanceName: []string{instanceName}})
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
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
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
		if errors.As(err, &e) {
			switch e.ExitStatus() {
			case 130:
				return nil
			}
		}
		return err
	}

	return nil
}

func NewServer(target string, addr string, multipassCertFilePath string, multipassKeyFilePath string, infoCh chan *Info) error {
	certPEMBlock, err := os.ReadFile(multipassCertFilePath)
	if err != nil {
		return err
	}
	keyPEMBlock, err := os.ReadFile(multipassKeyFilePath)
	if err != nil {
		return err
	}

	multipassCertificate, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return err
	}

	multipassTransportCredentials := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{multipassCertificate},
		InsecureSkipVerify: true,
	})

	var dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(multipassTransportCredentials)}
	conn, err := grpc.NewClient(target, dialOpts...)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:0", addr))
	if err != nil {
		return err
	}

	infoCh <- &Info{
		Port: lis.Addr().(*net.TCPAddr).Port,
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterRpcServer(grpcServer, &server{
		multipassClient: multipass.NewRpcClient(conn),
	})
	return grpcServer.Serve(lis)
}
