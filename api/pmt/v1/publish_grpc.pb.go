// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: pmt/v1/publish.proto

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
	Publish_Pub_FullMethodName = "/api.pmt.v1.Publish/Pub"
)

// PublishClient is the client API for Publish service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublishClient interface {
	Pub(ctx context.Context, in *PubRequest, opts ...grpc.CallOption) (*PubReply, error)
}

type publishClient struct {
	cc grpc.ClientConnInterface
}

func NewPublishClient(cc grpc.ClientConnInterface) PublishClient {
	return &publishClient{cc}
}

func (c *publishClient) Pub(ctx context.Context, in *PubRequest, opts ...grpc.CallOption) (*PubReply, error) {
	out := new(PubReply)
	err := c.cc.Invoke(ctx, Publish_Pub_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublishServer is the server API for Publish service.
// All implementations must embed UnimplementedPublishServer
// for forward compatibility
type PublishServer interface {
	Pub(context.Context, *PubRequest) (*PubReply, error)
	mustEmbedUnimplementedPublishServer()
}

// UnimplementedPublishServer must be embedded to have forward compatible implementations.
type UnimplementedPublishServer struct {
}

func (UnimplementedPublishServer) Pub(context.Context, *PubRequest) (*PubReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pub not implemented")
}
func (UnimplementedPublishServer) mustEmbedUnimplementedPublishServer() {}

// UnsafePublishServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublishServer will
// result in compilation errors.
type UnsafePublishServer interface {
	mustEmbedUnimplementedPublishServer()
}

func RegisterPublishServer(s grpc.ServiceRegistrar, srv PublishServer) {
	s.RegisterService(&Publish_ServiceDesc, srv)
}

func _Publish_Pub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).Pub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Publish_Pub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).Pub(ctx, req.(*PubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Publish_ServiceDesc is the grpc.ServiceDesc for Publish service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Publish_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.pmt.v1.Publish",
	HandlerType: (*PublishServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pub",
			Handler:    _Publish_Pub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pmt/v1/publish.proto",
}