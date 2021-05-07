// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package zookeeper

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

// ZooKeeperClient is the client API for ZooKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZooKeeperClient interface {
	// 定义SayHello方法
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	UserCheckIn(ctx context.Context, in *UserCheckInRequest, opts ...grpc.CallOption) (*UserCheckInResponse, error)
	UserCheckOut(ctx context.Context, in *UserCheckOutRequest, opts ...grpc.CallOption) (*UserCheckOutResponse, error)
	FindUser(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserResponse, error)
}

type zooKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewZooKeeperClient(cc grpc.ClientConnInterface) ZooKeeperClient {
	return &zooKeeperClient{cc}
}

func (c *zooKeeperClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/zookeeper.ZooKeeper/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zooKeeperClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/zookeeper.ZooKeeper/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zooKeeperClient) UserCheckIn(ctx context.Context, in *UserCheckInRequest, opts ...grpc.CallOption) (*UserCheckInResponse, error) {
	out := new(UserCheckInResponse)
	err := c.cc.Invoke(ctx, "/zookeeper.ZooKeeper/UserCheckIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zooKeeperClient) UserCheckOut(ctx context.Context, in *UserCheckOutRequest, opts ...grpc.CallOption) (*UserCheckOutResponse, error) {
	out := new(UserCheckOutResponse)
	err := c.cc.Invoke(ctx, "/zookeeper.ZooKeeper/UserCheckOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zooKeeperClient) FindUser(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserResponse, error) {
	out := new(FindUserResponse)
	err := c.cc.Invoke(ctx, "/zookeeper.ZooKeeper/FindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZooKeeperServer is the server API for ZooKeeper service.
// All implementations must embed UnimplementedZooKeeperServer
// for forward compatibility
type ZooKeeperServer interface {
	// 定义SayHello方法
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	UserCheckIn(context.Context, *UserCheckInRequest) (*UserCheckInResponse, error)
	UserCheckOut(context.Context, *UserCheckOutRequest) (*UserCheckOutResponse, error)
	FindUser(context.Context, *FindUserRequest) (*FindUserResponse, error)
	mustEmbedUnimplementedZooKeeperServer()
}

// UnimplementedZooKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedZooKeeperServer struct {
}

func (UnimplementedZooKeeperServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedZooKeeperServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedZooKeeperServer) UserCheckIn(context.Context, *UserCheckInRequest) (*UserCheckInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCheckIn not implemented")
}
func (UnimplementedZooKeeperServer) UserCheckOut(context.Context, *UserCheckOutRequest) (*UserCheckOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCheckOut not implemented")
}
func (UnimplementedZooKeeperServer) FindUser(context.Context, *FindUserRequest) (*FindUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}
func (UnimplementedZooKeeperServer) mustEmbedUnimplementedZooKeeperServer() {}

// UnsafeZooKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZooKeeperServer will
// result in compilation errors.
type UnsafeZooKeeperServer interface {
	mustEmbedUnimplementedZooKeeperServer()
}

func RegisterZooKeeperServer(s grpc.ServiceRegistrar, srv ZooKeeperServer) {
	s.RegisterService(&ZooKeeper_ServiceDesc, srv)
}

func _ZooKeeper_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZooKeeperServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zookeeper.ZooKeeper/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZooKeeperServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZooKeeper_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZooKeeperServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zookeeper.ZooKeeper/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZooKeeperServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZooKeeper_UserCheckIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCheckInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZooKeeperServer).UserCheckIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zookeeper.ZooKeeper/UserCheckIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZooKeeperServer).UserCheckIn(ctx, req.(*UserCheckInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZooKeeper_UserCheckOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCheckOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZooKeeperServer).UserCheckOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zookeeper.ZooKeeper/UserCheckOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZooKeeperServer).UserCheckOut(ctx, req.(*UserCheckOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZooKeeper_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZooKeeperServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zookeeper.ZooKeeper/FindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZooKeeperServer).FindUser(ctx, req.(*FindUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ZooKeeper_ServiceDesc is the grpc.ServiceDesc for ZooKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZooKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zookeeper.ZooKeeper",
	HandlerType: (*ZooKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ZooKeeper_Register_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _ZooKeeper_Heartbeat_Handler,
		},
		{
			MethodName: "UserCheckIn",
			Handler:    _ZooKeeper_UserCheckIn_Handler,
		},
		{
			MethodName: "UserCheckOut",
			Handler:    _ZooKeeper_UserCheckOut_Handler,
		},
		{
			MethodName: "FindUser",
			Handler:    _ZooKeeper_FindUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/zookeeper.proto",
}