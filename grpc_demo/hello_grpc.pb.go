// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/hello.proto

package grpc_demo

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

// GretterClient is the client API for Gretter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GretterClient interface {
	GetGreeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error)
}

type gretterClient struct {
	cc grpc.ClientConnInterface
}

func NewGretterClient(cc grpc.ClientConnInterface) GretterClient {
	return &gretterClient{cc}
}

func (c *gretterClient) GetGreeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error) {
	out := new(GreetingResponse)
	err := c.cc.Invoke(ctx, "/grpc_demo.Gretter/GetGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GretterServer is the server API for Gretter service.
// All implementations must embed UnimplementedGretterServer
// for forward compatibility
type GretterServer interface {
	GetGreeting(context.Context, *GreetingRequest) (*GreetingResponse, error)
	mustEmbedUnimplementedGretterServer()
}

// UnimplementedGretterServer must be embedded to have forward compatible implementations.
type UnimplementedGretterServer struct {
}

func (UnimplementedGretterServer) GetGreeting(context.Context, *GreetingRequest) (*GreetingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreeting not implemented")
}
func (UnimplementedGretterServer) mustEmbedUnimplementedGretterServer() {}

// UnsafeGretterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GretterServer will
// result in compilation errors.
type UnsafeGretterServer interface {
	mustEmbedUnimplementedGretterServer()
}

func RegisterGretterServer(s grpc.ServiceRegistrar, srv GretterServer) {
	s.RegisterService(&Gretter_ServiceDesc, srv)
}

func _Gretter_GetGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GretterServer).GetGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_demo.Gretter/GetGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GretterServer).GetGreeting(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gretter_ServiceDesc is the grpc.ServiceDesc for Gretter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gretter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_demo.Gretter",
	HandlerType: (*GretterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGreeting",
			Handler:    _Gretter_GetGreeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello.proto",
}
