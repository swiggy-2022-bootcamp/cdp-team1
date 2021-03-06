// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: productPriceService.proto

package productPrice

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

// ProductPriceServiceClient is the client API for ProductPriceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductPriceServiceClient interface {
	GetTotalPriceForProducts(ctx context.Context, in *ProductsPriceRequests, opts ...grpc.CallOption) (*ResponsePrice, error)
}

type productPriceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductPriceServiceClient(cc grpc.ClientConnInterface) ProductPriceServiceClient {
	return &productPriceServiceClient{cc}
}

func (c *productPriceServiceClient) GetTotalPriceForProducts(ctx context.Context, in *ProductsPriceRequests, opts ...grpc.CallOption) (*ResponsePrice, error) {
	out := new(ResponsePrice)
	err := c.cc.Invoke(ctx, "/proto.ProductPriceService/getTotalPriceForProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductPriceServiceServer is the server API for ProductPriceService service.
// All implementations must embed UnimplementedProductPriceServiceServer
// for forward compatibility
type ProductPriceServiceServer interface {
	GetTotalPriceForProducts(context.Context, *ProductsPriceRequests) (*ResponsePrice, error)
	mustEmbedUnimplementedProductPriceServiceServer()
}

// UnimplementedProductPriceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductPriceServiceServer struct {
}

func (UnimplementedProductPriceServiceServer) GetTotalPriceForProducts(context.Context, *ProductsPriceRequests) (*ResponsePrice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalPriceForProducts not implemented")
}
func (UnimplementedProductPriceServiceServer) mustEmbedUnimplementedProductPriceServiceServer() {}

// UnsafeProductPriceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductPriceServiceServer will
// result in compilation errors.
type UnsafeProductPriceServiceServer interface {
	mustEmbedUnimplementedProductPriceServiceServer()
}

func RegisterProductPriceServiceServer(s grpc.ServiceRegistrar, srv ProductPriceServiceServer) {
	s.RegisterService(&ProductPriceService_ServiceDesc, srv)
}

func _ProductPriceService_GetTotalPriceForProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsPriceRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductPriceServiceServer).GetTotalPriceForProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductPriceService/getTotalPriceForProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductPriceServiceServer).GetTotalPriceForProducts(ctx, req.(*ProductsPriceRequests))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductPriceService_ServiceDesc is the grpc.ServiceDesc for ProductPriceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductPriceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductPriceService",
	HandlerType: (*ProductPriceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getTotalPriceForProducts",
			Handler:    _ProductPriceService_GetTotalPriceForProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productPriceService.proto",
}
