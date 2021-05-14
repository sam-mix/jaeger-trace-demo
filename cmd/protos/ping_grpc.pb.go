// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RpcServerClient is the client API for RpcServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcServerClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error)
}

type rpcServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcServerClient(cc grpc.ClientConnInterface) RpcServerClient {
	return &rpcServerClient{cc}
}

func (c *rpcServerClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error) {
	out := new(PingRes)
	err := c.cc.Invoke(ctx, "/protos.rpcServer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServerServer is the server API for RpcServer service.
// All implementations must embed UnimplementedRpcServerServer
// for forward compatibility
type RpcServerServer interface {
	Ping(context.Context, *PingReq) (*PingRes, error)
	mustEmbedUnimplementedRpcServerServer()
}

// UnimplementedRpcServerServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServerServer struct {
}

func (UnimplementedRpcServerServer) Ping(context.Context, *PingReq) (*PingRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedRpcServerServer) mustEmbedUnimplementedRpcServerServer() {}

// UnsafeRpcServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServerServer will
// result in compilation errors.
type UnsafeRpcServerServer interface {
	mustEmbedUnimplementedRpcServerServer()
}

func RegisterRpcServerServer(s grpc.ServiceRegistrar, srv RpcServerServer) {
	s.RegisterService(&RpcServer_ServiceDesc, srv)
}

func _RpcServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.rpcServer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServerServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcServer_ServiceDesc is the grpc.ServiceDesc for RpcServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.rpcServer",
	HandlerType: (*RpcServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _RpcServer_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}
