// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: proto/faucet.proto

package faucet

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

// FaucetServiceClient is the client API for FaucetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FaucetServiceClient interface {
	VerifyTweet(ctx context.Context, in *VerifyTweetRequest, opts ...grpc.CallOption) (*VerifyTweetResponse, error)
	GetLinkedTwitters(ctx context.Context, in *GetLinkedTwittersRequest, opts ...grpc.CallOption) (*GetLinkedTwittersResponse, error)
	GetLinkedTwitterForAddress(ctx context.Context, in *LinkedTwitterForAddressRequest, opts ...grpc.CallOption) (*LinkedTwitterForAddressResponse, error)
	GetLinkedAddressForTwitter(ctx context.Context, in *LinkedAddressForTwitterRequest, opts ...grpc.CallOption) (*LinkedAddressForTwitterResponse, error)
}

type faucetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFaucetServiceClient(cc grpc.ClientConnInterface) FaucetServiceClient {
	return &faucetServiceClient{cc}
}

func (c *faucetServiceClient) VerifyTweet(ctx context.Context, in *VerifyTweetRequest, opts ...grpc.CallOption) (*VerifyTweetResponse, error) {
	out := new(VerifyTweetResponse)
	err := c.cc.Invoke(ctx, "/faucet.FaucetService/VerifyTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faucetServiceClient) GetLinkedTwitters(ctx context.Context, in *GetLinkedTwittersRequest, opts ...grpc.CallOption) (*GetLinkedTwittersResponse, error) {
	out := new(GetLinkedTwittersResponse)
	err := c.cc.Invoke(ctx, "/faucet.FaucetService/GetLinkedTwitters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faucetServiceClient) GetLinkedTwitterForAddress(ctx context.Context, in *LinkedTwitterForAddressRequest, opts ...grpc.CallOption) (*LinkedTwitterForAddressResponse, error) {
	out := new(LinkedTwitterForAddressResponse)
	err := c.cc.Invoke(ctx, "/faucet.FaucetService/GetLinkedTwitterForAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faucetServiceClient) GetLinkedAddressForTwitter(ctx context.Context, in *LinkedAddressForTwitterRequest, opts ...grpc.CallOption) (*LinkedAddressForTwitterResponse, error) {
	out := new(LinkedAddressForTwitterResponse)
	err := c.cc.Invoke(ctx, "/faucet.FaucetService/GetLinkedAddressForTwitter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaucetServiceServer is the server API for FaucetService service.
// All implementations must embed UnimplementedFaucetServiceServer
// for forward compatibility
type FaucetServiceServer interface {
	VerifyTweet(context.Context, *VerifyTweetRequest) (*VerifyTweetResponse, error)
	GetLinkedTwitters(context.Context, *GetLinkedTwittersRequest) (*GetLinkedTwittersResponse, error)
	GetLinkedTwitterForAddress(context.Context, *LinkedTwitterForAddressRequest) (*LinkedTwitterForAddressResponse, error)
	GetLinkedAddressForTwitter(context.Context, *LinkedAddressForTwitterRequest) (*LinkedAddressForTwitterResponse, error)
	mustEmbedUnimplementedFaucetServiceServer()
}

// UnimplementedFaucetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFaucetServiceServer struct {
}

func (UnimplementedFaucetServiceServer) VerifyTweet(context.Context, *VerifyTweetRequest) (*VerifyTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTweet not implemented")
}
func (UnimplementedFaucetServiceServer) GetLinkedTwitters(context.Context, *GetLinkedTwittersRequest) (*GetLinkedTwittersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkedTwitters not implemented")
}
func (UnimplementedFaucetServiceServer) GetLinkedTwitterForAddress(context.Context, *LinkedTwitterForAddressRequest) (*LinkedTwitterForAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkedTwitterForAddress not implemented")
}
func (UnimplementedFaucetServiceServer) GetLinkedAddressForTwitter(context.Context, *LinkedAddressForTwitterRequest) (*LinkedAddressForTwitterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkedAddressForTwitter not implemented")
}
func (UnimplementedFaucetServiceServer) mustEmbedUnimplementedFaucetServiceServer() {}

// UnsafeFaucetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FaucetServiceServer will
// result in compilation errors.
type UnsafeFaucetServiceServer interface {
	mustEmbedUnimplementedFaucetServiceServer()
}

func RegisterFaucetServiceServer(s grpc.ServiceRegistrar, srv FaucetServiceServer) {
	s.RegisterService(&FaucetService_ServiceDesc, srv)
}

func _FaucetService_VerifyTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaucetServiceServer).VerifyTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faucet.FaucetService/VerifyTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaucetServiceServer).VerifyTweet(ctx, req.(*VerifyTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaucetService_GetLinkedTwitters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkedTwittersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaucetServiceServer).GetLinkedTwitters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faucet.FaucetService/GetLinkedTwitters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaucetServiceServer).GetLinkedTwitters(ctx, req.(*GetLinkedTwittersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaucetService_GetLinkedTwitterForAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkedTwitterForAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaucetServiceServer).GetLinkedTwitterForAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faucet.FaucetService/GetLinkedTwitterForAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaucetServiceServer).GetLinkedTwitterForAddress(ctx, req.(*LinkedTwitterForAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaucetService_GetLinkedAddressForTwitter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkedAddressForTwitterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaucetServiceServer).GetLinkedAddressForTwitter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faucet.FaucetService/GetLinkedAddressForTwitter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaucetServiceServer).GetLinkedAddressForTwitter(ctx, req.(*LinkedAddressForTwitterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FaucetService_ServiceDesc is the grpc.ServiceDesc for FaucetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FaucetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "faucet.FaucetService",
	HandlerType: (*FaucetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyTweet",
			Handler:    _FaucetService_VerifyTweet_Handler,
		},
		{
			MethodName: "GetLinkedTwitters",
			Handler:    _FaucetService_GetLinkedTwitters_Handler,
		},
		{
			MethodName: "GetLinkedTwitterForAddress",
			Handler:    _FaucetService_GetLinkedTwitterForAddress_Handler,
		},
		{
			MethodName: "GetLinkedAddressForTwitter",
			Handler:    _FaucetService_GetLinkedAddressForTwitter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/faucet.proto",
}