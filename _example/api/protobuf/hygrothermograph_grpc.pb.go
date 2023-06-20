// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.2
// source: hygrothermograph.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	HygrothermographService_GetHygrothermograph_FullMethodName = "/protobuf.api.HygrothermographService/GetHygrothermograph"
)

// HygrothermographServiceClient is the client API for HygrothermographService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HygrothermographServiceClient interface {
	GetHygrothermograph(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Hygrothermograph, error)
}

type hygrothermographServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHygrothermographServiceClient(cc grpc.ClientConnInterface) HygrothermographServiceClient {
	return &hygrothermographServiceClient{cc}
}

func (c *hygrothermographServiceClient) GetHygrothermograph(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Hygrothermograph, error) {
	out := new(Hygrothermograph)
	err := c.cc.Invoke(ctx, HygrothermographService_GetHygrothermograph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HygrothermographServiceServer is the server API for HygrothermographService service.
// All implementations must embed UnimplementedHygrothermographServiceServer
// for forward compatibility
type HygrothermographServiceServer interface {
	GetHygrothermograph(context.Context, *emptypb.Empty) (*Hygrothermograph, error)
	mustEmbedUnimplementedHygrothermographServiceServer()
}

// UnimplementedHygrothermographServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHygrothermographServiceServer struct {
}

func (UnimplementedHygrothermographServiceServer) GetHygrothermograph(context.Context, *emptypb.Empty) (*Hygrothermograph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHygrothermograph not implemented")
}
func (UnimplementedHygrothermographServiceServer) mustEmbedUnimplementedHygrothermographServiceServer() {
}

// UnsafeHygrothermographServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HygrothermographServiceServer will
// result in compilation errors.
type UnsafeHygrothermographServiceServer interface {
	mustEmbedUnimplementedHygrothermographServiceServer()
}

func RegisterHygrothermographServiceServer(s grpc.ServiceRegistrar, srv HygrothermographServiceServer) {
	s.RegisterService(&HygrothermographService_ServiceDesc, srv)
}

func _HygrothermographService_GetHygrothermograph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HygrothermographServiceServer).GetHygrothermograph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HygrothermographService_GetHygrothermograph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HygrothermographServiceServer).GetHygrothermograph(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HygrothermographService_ServiceDesc is the grpc.ServiceDesc for HygrothermographService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HygrothermographService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.api.HygrothermographService",
	HandlerType: (*HygrothermographServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHygrothermograph",
			Handler:    _HygrothermographService_GetHygrothermograph_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hygrothermograph.proto",
}
