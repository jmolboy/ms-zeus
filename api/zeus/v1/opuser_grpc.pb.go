// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: zeus/v1/opuser.proto

package v1

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

const (
	OpUser_Current_FullMethodName = "/api.zeus.v1.OpUser/Current"
	OpUser_Logout_FullMethodName  = "/api.zeus.v1.OpUser/Logout"
)

// OpUserClient is the client API for OpUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpUserClient interface {
	Current(ctx context.Context, in *CurrentRequest, opts ...grpc.CallOption) (*CurrentReply, error)
	Logout(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutReply, error)
}

type opUserClient struct {
	cc grpc.ClientConnInterface
}

func NewOpUserClient(cc grpc.ClientConnInterface) OpUserClient {
	return &opUserClient{cc}
}

func (c *opUserClient) Current(ctx context.Context, in *CurrentRequest, opts ...grpc.CallOption) (*CurrentReply, error) {
	out := new(CurrentReply)
	err := c.cc.Invoke(ctx, OpUser_Current_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opUserClient) Logout(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutReply, error) {
	out := new(LogOutReply)
	err := c.cc.Invoke(ctx, OpUser_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpUserServer is the server API for OpUser service.
// All implementations must embed UnimplementedOpUserServer
// for forward compatibility
type OpUserServer interface {
	Current(context.Context, *CurrentRequest) (*CurrentReply, error)
	Logout(context.Context, *LogOutRequest) (*LogOutReply, error)
	mustEmbedUnimplementedOpUserServer()
}

// UnimplementedOpUserServer must be embedded to have forward compatible implementations.
type UnimplementedOpUserServer struct {
}

func (UnimplementedOpUserServer) Current(context.Context, *CurrentRequest) (*CurrentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Current not implemented")
}
func (UnimplementedOpUserServer) Logout(context.Context, *LogOutRequest) (*LogOutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedOpUserServer) mustEmbedUnimplementedOpUserServer() {}

// UnsafeOpUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpUserServer will
// result in compilation errors.
type UnsafeOpUserServer interface {
	mustEmbedUnimplementedOpUserServer()
}

func RegisterOpUserServer(s grpc.ServiceRegistrar, srv OpUserServer) {
	s.RegisterService(&OpUser_ServiceDesc, srv)
}

func _OpUser_Current_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpUserServer).Current(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpUser_Current_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpUserServer).Current(ctx, req.(*CurrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpUser_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpUserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpUser_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpUserServer).Logout(ctx, req.(*LogOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OpUser_ServiceDesc is the grpc.ServiceDesc for OpUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.zeus.v1.OpUser",
	HandlerType: (*OpUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Current",
			Handler:    _OpUser_Current_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _OpUser_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zeus/v1/opuser.proto",
}
