package agent

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"multiverse/common"
	"multiverse/multipass"
	"net"
	"strconv"
	"sync"

	"github.com/google/uuid"

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
	sshMap          map[string]SSH
	sshMu           sync.RWMutex
}

func (s *server) GetSSH(uid string) SSH {
	s.sshMu.Lock()
	defer s.sshMu.Unlock()
	return s.sshMap[uid]
}

func (s *server) addSSH(uid string, ssh SSH) {
	s.sshMu.Lock()
	defer s.sshMu.Unlock()
	s.sshMap[uid] = ssh
}

func (s *server) removeSSH(uid string) {
	s.sshMu.Lock()
	defer s.sshMu.Unlock()
	if ssh, ok := s.sshMap[uid]; ok {
		err := ssh.Close()
		if err != nil {
			log.Printf("failed to close ssh: %v", err)
		}
	}
	delete(s.sshMap, uid)
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

	res, err := common.ExecuteOnceWithBidiClient(stream, &multipass.ListRequest{})
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

type windowSize struct {
	sig    chan *windowSize
	width  int64
	height int64
}

func (s *windowSize) setIfChanged(req *ShellRequest) {
	width := req.GetWidth()
	height := req.GetHeight()
	if s.width != width || s.height != height {
		s.width = width
		s.height = height
		s.sig <- s
	}
}

type shellRequestReader struct {
	stream     grpc.BidiStreamingServer[ShellRequest, ShellReply]
	windowSize *windowSize
}

func NewShellRequestReader(
	stream grpc.BidiStreamingServer[ShellRequest, ShellReply],
	height int, width int,
) *shellRequestReader {
	return &shellRequestReader{
		stream: stream,
		windowSize: &windowSize{
			width:  int64(width),
			height: int64(height),
			sig:    make(chan *windowSize, 1),
		},
	}
}

func (s *shellRequestReader) Read(p []byte) (n int, err error) {
	in, err := s.stream.Recv()
	if err != nil {
		close(s.windowSize.sig)
		return 0, err
	}
	s.windowSize.setIfChanged(in)
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

func (s *server) Shell(stream grpc.BidiStreamingServer[ShellRequest, ShellReply]) error { // nolint:funlen
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("metadata not found in context")
	}

	height := md.Get("height")
	if len(height) == 0 {
		return fmt.Errorf("height not found in context")
	}
	h, _ := strconv.Atoi(height[0])

	width := md.Get("width")
	if len(width) == 0 {
		return fmt.Errorf("width not found in context")
	}
	w, _ := strconv.Atoi(width[0])

	stdout := &shellReplyWriter{stream: stream}
	stderr := &shellReplyWriter{stream: stream, isErr: true}
	stdin := NewShellRequestReader(stream, h, w)

	sshInfoStream, err := s.multipassClient.SshInfo(context.Background())
	if err != nil {
		return err
	}

	instanceName := md.Get("instanceName")
	if len(instanceName) == 0 {
		return fmt.Errorf("instance name not found in context")
	}

	res, err := common.ExecuteOnceWithBidiClient(
		sshInfoStream,
		&multipass.SSHInfoRequest{
			InstanceName: []string{instanceName[0]},
		},
	)
	if err != nil {
		return err
	}

	info, ok := res.SshInfo[instanceName[0]]
	if !ok {
		return fmt.Errorf("instance not found: %s", instanceName)
	}

	id := uuid.Must(uuid.NewRandom()).String()
	defer log.Printf("ssh disconnected: %s", id)
	defer s.removeSSH(id)
	ssh := NewSSH(info.Host, int(info.Port), info.Username, []byte(info.PrivKeyBase64), stdout, stderr, stdin, h, w)
	s.addSSH(id, ssh)
	log.Printf("ssh connected: %s", id)
	go ssh.InheritSize(stdin.windowSize.sig)
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
		sshMap:          make(map[string]SSH),
		sshMu:           sync.RWMutex{},
	}
	RegisterRpcServer(grpcServer, server)
	return server, nil
}
