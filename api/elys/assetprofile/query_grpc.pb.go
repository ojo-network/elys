// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: elys/assetprofile/query.proto

package assetprofile

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
	Query_Params_FullMethodName       = "/elys.assetprofile.Query/Params"
	Query_Entry_FullMethodName        = "/elys.assetprofile.Query/Entry"
	Query_EntryByDenom_FullMethodName = "/elys.assetprofile.Query/EntryByDenom"
	Query_EntryAll_FullMethodName     = "/elys.assetprofile.Query/EntryAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Entry items.
	Entry(ctx context.Context, in *QueryEntryRequest, opts ...grpc.CallOption) (*QueryEntryResponse, error)
	EntryByDenom(ctx context.Context, in *QueryEntryByDenomRequest, opts ...grpc.CallOption) (*QueryEntryByDenomResponse, error)
	EntryAll(ctx context.Context, in *QueryAllEntryRequest, opts ...grpc.CallOption) (*QueryAllEntryResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Entry(ctx context.Context, in *QueryEntryRequest, opts ...grpc.CallOption) (*QueryEntryResponse, error) {
	out := new(QueryEntryResponse)
	err := c.cc.Invoke(ctx, Query_Entry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EntryByDenom(ctx context.Context, in *QueryEntryByDenomRequest, opts ...grpc.CallOption) (*QueryEntryByDenomResponse, error) {
	out := new(QueryEntryByDenomResponse)
	err := c.cc.Invoke(ctx, Query_EntryByDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EntryAll(ctx context.Context, in *QueryAllEntryRequest, opts ...grpc.CallOption) (*QueryAllEntryResponse, error) {
	out := new(QueryAllEntryResponse)
	err := c.cc.Invoke(ctx, Query_EntryAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Entry items.
	Entry(context.Context, *QueryEntryRequest) (*QueryEntryResponse, error)
	EntryByDenom(context.Context, *QueryEntryByDenomRequest) (*QueryEntryByDenomResponse, error)
	EntryAll(context.Context, *QueryAllEntryRequest) (*QueryAllEntryResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Entry(context.Context, *QueryEntryRequest) (*QueryEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Entry not implemented")
}
func (UnimplementedQueryServer) EntryByDenom(context.Context, *QueryEntryByDenomRequest) (*QueryEntryByDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntryByDenom not implemented")
}
func (UnimplementedQueryServer) EntryAll(context.Context, *QueryAllEntryRequest) (*QueryAllEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntryAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Entry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Entry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Entry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Entry(ctx, req.(*QueryEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EntryByDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEntryByDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EntryByDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_EntryByDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EntryByDenom(ctx, req.(*QueryEntryByDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EntryAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EntryAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_EntryAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EntryAll(ctx, req.(*QueryAllEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "elys.assetprofile.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Entry",
			Handler:    _Query_Entry_Handler,
		},
		{
			MethodName: "EntryByDenom",
			Handler:    _Query_EntryByDenom_Handler,
		},
		{
			MethodName: "EntryAll",
			Handler:    _Query_EntryAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elys/assetprofile/query.proto",
}
