// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: agent/agent.proto

package agent

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Rpc_List_FullMethodName  = "/agent.Rpc/list"
	Rpc_Shell_FullMethodName = "/agent.Rpc/shell"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Shell(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ShellRequest, ShellReply], error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListReply)
	err := c.cc.Invoke(ctx, Rpc_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Shell(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ShellRequest, ShellReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Rpc_ServiceDesc.Streams[0], Rpc_Shell_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ShellRequest, ShellReply]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Rpc_ShellClient = grpc.BidiStreamingClient[ShellRequest, ShellReply]

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility.
type RpcServer interface {
	List(context.Context, *ListRequest) (*ListReply, error)
	Shell(grpc.BidiStreamingServer[ShellRequest, ShellReply]) error
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRpcServer struct{}

func (UnimplementedRpcServer) List(context.Context, *ListRequest) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRpcServer) Shell(grpc.BidiStreamingServer[ShellRequest, ShellReply]) error {
	return status.Errorf(codes.Unimplemented, "method Shell not implemented")
}
func (UnimplementedRpcServer) mustEmbedUnimplementedRpcServer() {}
func (UnimplementedRpcServer) testEmbeddedByValue()             {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	// If the following call pancis, it indicates UnimplementedRpcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Shell_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcServer).Shell(&grpc.GenericServerStream[ShellRequest, ShellReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Rpc_ShellServer = grpc.BidiStreamingServer[ShellRequest, ShellReply]

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _Rpc_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "shell",
			Handler:       _Rpc_Shell_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent/agent.proto",
}
