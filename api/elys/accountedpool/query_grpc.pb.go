// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: elys/accountedpool/query.proto

package accountedpool

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
	Query_AccountedPool_FullMethodName    = "/elys.accountedpool.Query/AccountedPool"
	Query_AccountedPoolAll_FullMethodName = "/elys.accountedpool.Query/AccountedPoolAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a list of AccountedPool items.
	AccountedPool(ctx context.Context, in *QueryGetAccountedPoolRequest, opts ...grpc.CallOption) (*QueryGetAccountedPoolResponse, error)
	AccountedPoolAll(ctx context.Context, in *QueryAllAccountedPoolRequest, opts ...grpc.CallOption) (*QueryAllAccountedPoolResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AccountedPool(ctx context.Context, in *QueryGetAccountedPoolRequest, opts ...grpc.CallOption) (*QueryGetAccountedPoolResponse, error) {
	out := new(QueryGetAccountedPoolResponse)
	err := c.cc.Invoke(ctx, Query_AccountedPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AccountedPoolAll(ctx context.Context, in *QueryAllAccountedPoolRequest, opts ...grpc.CallOption) (*QueryAllAccountedPoolResponse, error) {
	out := new(QueryAllAccountedPoolResponse)
	err := c.cc.Invoke(ctx, Query_AccountedPoolAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Queries a list of AccountedPool items.
	AccountedPool(context.Context, *QueryGetAccountedPoolRequest) (*QueryGetAccountedPoolResponse, error)
	AccountedPoolAll(context.Context, *QueryAllAccountedPoolRequest) (*QueryAllAccountedPoolResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) AccountedPool(context.Context, *QueryGetAccountedPoolRequest) (*QueryGetAccountedPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountedPool not implemented")
}
func (UnimplementedQueryServer) AccountedPoolAll(context.Context, *QueryAllAccountedPoolRequest) (*QueryAllAccountedPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountedPoolAll not implemented")
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

func _Query_AccountedPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAccountedPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AccountedPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AccountedPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AccountedPool(ctx, req.(*QueryGetAccountedPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AccountedPoolAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllAccountedPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AccountedPoolAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AccountedPoolAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AccountedPoolAll(ctx, req.(*QueryAllAccountedPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "elys.accountedpool.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountedPool",
			Handler:    _Query_AccountedPool_Handler,
		},
		{
			MethodName: "AccountedPoolAll",
			Handler:    _Query_AccountedPoolAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elys/accountedpool/query.proto",
}
