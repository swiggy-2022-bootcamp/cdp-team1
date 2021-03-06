// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service.proto

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

// QuantityServiceClient is the client API for QuantityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuantityServiceClient interface {
	GetQuantity(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type quantityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuantityServiceClient(cc grpc.ClientConnInterface) QuantityServiceClient {
	return &quantityServiceClient{cc}
}

func (c *quantityServiceClient) GetQuantity(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.QuantityService/getQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuantityServiceServer is the server API for QuantityService service.
// All implementations must embed UnimplementedQuantityServiceServer
// for forward compatibility
type QuantityServiceServer interface {
	GetQuantity(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedQuantityServiceServer()
}

// UnimplementedQuantityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuantityServiceServer struct {
}

func (UnimplementedQuantityServiceServer) GetQuantity(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuantity not implemented")
}
func (UnimplementedQuantityServiceServer) mustEmbedUnimplementedQuantityServiceServer() {}

// UnsafeQuantityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuantityServiceServer will
// result in compilation errors.
type UnsafeQuantityServiceServer interface {
	mustEmbedUnimplementedQuantityServiceServer()
}

func RegisterQuantityServiceServer(s grpc.ServiceRegistrar, srv QuantityServiceServer) {
	s.RegisterService(&QuantityService_ServiceDesc, srv)
}

func _QuantityService_GetQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuantityServiceServer).GetQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuantityService/getQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuantityServiceServer).GetQuantity(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// QuantityService_ServiceDesc is the grpc.ServiceDesc for QuantityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuantityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.QuantityService",
	HandlerType: (*QuantityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getQuantity",
			Handler:    _QuantityService_GetQuantity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
