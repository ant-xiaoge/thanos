// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.0
// source: metadata/metadatapb/rpc.proto

package metadatapb

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
	Metadata_MetricMetadata_FullMethodName = "/thanos.Metadata/MetricMetadata"
)

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataClient interface {
	MetricMetadata(ctx context.Context, in *MetricMetadataRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MetricMetadataResponse], error)
}

type metadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataClient(cc grpc.ClientConnInterface) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) MetricMetadata(ctx context.Context, in *MetricMetadataRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MetricMetadataResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Metadata_ServiceDesc.Streams[0], Metadata_MetricMetadata_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MetricMetadataRequest, MetricMetadataResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Metadata_MetricMetadataClient = grpc.ServerStreamingClient[MetricMetadataResponse]

// MetadataServer is the server API for Metadata service.
// All implementations must embed UnimplementedMetadataServer
// for forward compatibility.
type MetadataServer interface {
	MetricMetadata(*MetricMetadataRequest, grpc.ServerStreamingServer[MetricMetadataResponse]) error
	mustEmbedUnimplementedMetadataServer()
}

// UnimplementedMetadataServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMetadataServer struct{}

func (UnimplementedMetadataServer) MetricMetadata(*MetricMetadataRequest, grpc.ServerStreamingServer[MetricMetadataResponse]) error {
	return status.Errorf(codes.Unimplemented, "method MetricMetadata not implemented")
}
func (UnimplementedMetadataServer) mustEmbedUnimplementedMetadataServer() {}
func (UnimplementedMetadataServer) testEmbeddedByValue()                  {}

// UnsafeMetadataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServer will
// result in compilation errors.
type UnsafeMetadataServer interface {
	mustEmbedUnimplementedMetadataServer()
}

func RegisterMetadataServer(s grpc.ServiceRegistrar, srv MetadataServer) {
	// If the following call pancis, it indicates UnimplementedMetadataServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Metadata_ServiceDesc, srv)
}

func _Metadata_MetricMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetricMetadataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetadataServer).MetricMetadata(m, &grpc.GenericServerStream[MetricMetadataRequest, MetricMetadataResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Metadata_MetricMetadataServer = grpc.ServerStreamingServer[MetricMetadataResponse]

// Metadata_ServiceDesc is the grpc.ServiceDesc for Metadata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metadata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thanos.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MetricMetadata",
			Handler:       _Metadata_MetricMetadata_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "metadata/metadatapb/rpc.proto",
}