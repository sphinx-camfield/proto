// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: svc/lk-token-provider.proto

package svc

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
	LivekitTokenProvider_Request_FullMethodName = "/camfield.svc.LivekitTokenProvider/Request"
)

// LivekitTokenProviderClient is the client API for LivekitTokenProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LivekitTokenProviderClient interface {
	Request(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RoomTokenReply, error)
}

type livekitTokenProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewLivekitTokenProviderClient(cc grpc.ClientConnInterface) LivekitTokenProviderClient {
	return &livekitTokenProviderClient{cc}
}

func (c *livekitTokenProviderClient) Request(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RoomTokenReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoomTokenReply)
	err := c.cc.Invoke(ctx, LivekitTokenProvider_Request_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LivekitTokenProviderServer is the server API for LivekitTokenProvider service.
// All implementations must embed UnimplementedLivekitTokenProviderServer
// for forward compatibility.
type LivekitTokenProviderServer interface {
	Request(context.Context, *EmptyRequest) (*RoomTokenReply, error)
	mustEmbedUnimplementedLivekitTokenProviderServer()
}

// UnimplementedLivekitTokenProviderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLivekitTokenProviderServer struct{}

func (UnimplementedLivekitTokenProviderServer) Request(context.Context, *EmptyRequest) (*RoomTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedLivekitTokenProviderServer) mustEmbedUnimplementedLivekitTokenProviderServer() {}
func (UnimplementedLivekitTokenProviderServer) testEmbeddedByValue()                              {}

// UnsafeLivekitTokenProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LivekitTokenProviderServer will
// result in compilation errors.
type UnsafeLivekitTokenProviderServer interface {
	mustEmbedUnimplementedLivekitTokenProviderServer()
}

func RegisterLivekitTokenProviderServer(s grpc.ServiceRegistrar, srv LivekitTokenProviderServer) {
	// If the following call pancis, it indicates UnimplementedLivekitTokenProviderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LivekitTokenProvider_ServiceDesc, srv)
}

func _LivekitTokenProvider_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivekitTokenProviderServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivekitTokenProvider_Request_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivekitTokenProviderServer).Request(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LivekitTokenProvider_ServiceDesc is the grpc.ServiceDesc for LivekitTokenProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LivekitTokenProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "camfield.svc.LivekitTokenProvider",
	HandlerType: (*LivekitTokenProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _LivekitTokenProvider_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc/lk-token-provider.proto",
}
