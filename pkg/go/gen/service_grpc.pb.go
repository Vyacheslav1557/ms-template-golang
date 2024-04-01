// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: service.proto

package gen

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

// SampleServiceClient is the client API for SampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	Count(ctx context.Context, opts ...grpc.CallOption) (SampleService_CountClient, error)
}

type sampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleServiceClient(cc grpc.ClientConnInterface) SampleServiceClient {
	return &sampleServiceClient{cc}
}

func (c *sampleServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/sample.SampleService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleServiceClient) Count(ctx context.Context, opts ...grpc.CallOption) (SampleService_CountClient, error) {
	stream, err := c.cc.NewStream(ctx, &SampleService_ServiceDesc.Streams[0], "/sample.SampleService/Count", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleServiceCountClient{stream}
	return x, nil
}

type SampleService_CountClient interface {
	Send(*CountRequest) error
	Recv() (*CountResponse, error)
	grpc.ClientStream
}

type sampleServiceCountClient struct {
	grpc.ClientStream
}

func (x *sampleServiceCountClient) Send(m *CountRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleServiceCountClient) Recv() (*CountResponse, error) {
	m := new(CountResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SampleServiceServer is the server API for SampleService service.
// All implementations must embed UnimplementedSampleServiceServer
// for forward compatibility
type SampleServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	Count(SampleService_CountServer) error
	mustEmbedUnimplementedSampleServiceServer()
}

// UnimplementedSampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSampleServiceServer struct {
}

func (UnimplementedSampleServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedSampleServiceServer) Count(SampleService_CountServer) error {
	return status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedSampleServiceServer) mustEmbedUnimplementedSampleServiceServer() {}

// UnsafeSampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleServiceServer will
// result in compilation errors.
type UnsafeSampleServiceServer interface {
	mustEmbedUnimplementedSampleServiceServer()
}

func RegisterSampleServiceServer(s grpc.ServiceRegistrar, srv SampleServiceServer) {
	s.RegisterService(&SampleService_ServiceDesc, srv)
}

func _SampleService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.SampleService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleService_Count_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SampleServiceServer).Count(&sampleServiceCountServer{stream})
}

type SampleService_CountServer interface {
	Send(*CountResponse) error
	Recv() (*CountRequest, error)
	grpc.ServerStream
}

type sampleServiceCountServer struct {
	grpc.ServerStream
}

func (x *sampleServiceCountServer) Send(m *CountResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleServiceCountServer) Recv() (*CountRequest, error) {
	m := new(CountRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SampleService_ServiceDesc is the grpc.ServiceDesc for SampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample.SampleService",
	HandlerType: (*SampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SampleService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Count",
			Handler:       _SampleService_Count_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
