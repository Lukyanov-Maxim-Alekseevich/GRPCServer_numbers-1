// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: adder.proto

package api

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
	Number_YesNumber_FullMethodName = "/api.Number/YesNumber"
	Number_NoNumber_FullMethodName  = "/api.Number/NoNumber"
)

// NumberClient is the client API for Number service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NumberClient interface {
	YesNumber(ctx context.Context, in *YesRequest, opts ...grpc.CallOption) (*YesResponse, error)
	NoNumber(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (*NoResponse, error)
}

type numberClient struct {
	cc grpc.ClientConnInterface
}

func NewNumberClient(cc grpc.ClientConnInterface) NumberClient {
	return &numberClient{cc}
}

func (c *numberClient) YesNumber(ctx context.Context, in *YesRequest, opts ...grpc.CallOption) (*YesResponse, error) {
	out := new(YesResponse)
	err := c.cc.Invoke(ctx, Number_YesNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *numberClient) NoNumber(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, Number_NoNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NumberServer is the server API for Number service.
// All implementations must embed UnimplementedNumberServer
// for forward compatibility
type NumberServer interface {
	YesNumber(context.Context, *YesRequest) (*YesResponse, error)
	NoNumber(context.Context, *NoRequest) (*NoResponse, error)
	mustEmbedUnimplementedNumberServer()
}

// UnimplementedNumberServer must be embedded to have forward compatible implementations.
type UnimplementedNumberServer struct {
}

func (UnimplementedNumberServer) YesNumber(context.Context, *YesRequest) (*YesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method YesNumber not implemented")
}
func (UnimplementedNumberServer) NoNumber(context.Context, *NoRequest) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoNumber not implemented")
}
func (UnimplementedNumberServer) mustEmbedUnimplementedNumberServer() {}

// UnsafeNumberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NumberServer will
// result in compilation errors.
type UnsafeNumberServer interface {
	mustEmbedUnimplementedNumberServer()
}

func RegisterNumberServer(s grpc.ServiceRegistrar, srv NumberServer) {
	s.RegisterService(&Number_ServiceDesc, srv)
}

func _Number_YesNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberServer).YesNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Number_YesNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberServer).YesNumber(ctx, req.(*YesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Number_NoNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberServer).NoNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Number_NoNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberServer).NoNumber(ctx, req.(*NoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Number_ServiceDesc is the grpc.ServiceDesc for Number service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Number_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Number",
	HandlerType: (*NumberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "YesNumber",
			Handler:    _Number_YesNumber_Handler,
		},
		{
			MethodName: "NoNumber",
			Handler:    _Number_NoNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adder.proto",
}