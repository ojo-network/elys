// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: elys/oracle/query.proto

package oracle

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
	Query_Params_FullMethodName            = "/elys.oracle.Query/Params"
	Query_BandPriceResult_FullMethodName   = "/elys.oracle.Query/BandPriceResult"
	Query_LastBandRequestId_FullMethodName = "/elys.oracle.Query/LastBandRequestId"
	Query_AssetInfo_FullMethodName         = "/elys.oracle.Query/AssetInfo"
	Query_AssetInfoAll_FullMethodName      = "/elys.oracle.Query/AssetInfoAll"
	Query_Price_FullMethodName             = "/elys.oracle.Query/Price"
	Query_PriceAll_FullMethodName          = "/elys.oracle.Query/PriceAll"
	Query_PriceFeeder_FullMethodName       = "/elys.oracle.Query/PriceFeeder"
	Query_PriceFeederAll_FullMethodName    = "/elys.oracle.Query/PriceFeederAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// BandPriceResult defines a rpc handler method for MsgRequestBandPrice.
	BandPriceResult(ctx context.Context, in *QueryBandPriceRequest, opts ...grpc.CallOption) (*QueryBandPriceResponse, error)
	// LastBandRequestId query the last BandPrice result id
	LastBandRequestId(ctx context.Context, in *QueryLastBandRequestIdRequest, opts ...grpc.CallOption) (*QueryLastBandRequestIdResponse, error)
	// Queries a AssetInfo by denom.
	AssetInfo(ctx context.Context, in *QueryGetAssetInfoRequest, opts ...grpc.CallOption) (*QueryGetAssetInfoResponse, error)
	// Queries a list of AssetInfo items.
	AssetInfoAll(ctx context.Context, in *QueryAllAssetInfoRequest, opts ...grpc.CallOption) (*QueryAllAssetInfoResponse, error)
	// Queries a Price by asset.
	Price(ctx context.Context, in *QueryGetPriceRequest, opts ...grpc.CallOption) (*QueryGetPriceResponse, error)
	// Queries a list of Price items.
	PriceAll(ctx context.Context, in *QueryAllPriceRequest, opts ...grpc.CallOption) (*QueryAllPriceResponse, error)
	// Queries a PriceFeeder by feeder.
	PriceFeeder(ctx context.Context, in *QueryGetPriceFeederRequest, opts ...grpc.CallOption) (*QueryGetPriceFeederResponse, error)
	// Queries a list of PriceFeeder items.
	PriceFeederAll(ctx context.Context, in *QueryAllPriceFeederRequest, opts ...grpc.CallOption) (*QueryAllPriceFeederResponse, error)
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

func (c *queryClient) BandPriceResult(ctx context.Context, in *QueryBandPriceRequest, opts ...grpc.CallOption) (*QueryBandPriceResponse, error) {
	out := new(QueryBandPriceResponse)
	err := c.cc.Invoke(ctx, Query_BandPriceResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LastBandRequestId(ctx context.Context, in *QueryLastBandRequestIdRequest, opts ...grpc.CallOption) (*QueryLastBandRequestIdResponse, error) {
	out := new(QueryLastBandRequestIdResponse)
	err := c.cc.Invoke(ctx, Query_LastBandRequestId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AssetInfo(ctx context.Context, in *QueryGetAssetInfoRequest, opts ...grpc.CallOption) (*QueryGetAssetInfoResponse, error) {
	out := new(QueryGetAssetInfoResponse)
	err := c.cc.Invoke(ctx, Query_AssetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AssetInfoAll(ctx context.Context, in *QueryAllAssetInfoRequest, opts ...grpc.CallOption) (*QueryAllAssetInfoResponse, error) {
	out := new(QueryAllAssetInfoResponse)
	err := c.cc.Invoke(ctx, Query_AssetInfoAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Price(ctx context.Context, in *QueryGetPriceRequest, opts ...grpc.CallOption) (*QueryGetPriceResponse, error) {
	out := new(QueryGetPriceResponse)
	err := c.cc.Invoke(ctx, Query_Price_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PriceAll(ctx context.Context, in *QueryAllPriceRequest, opts ...grpc.CallOption) (*QueryAllPriceResponse, error) {
	out := new(QueryAllPriceResponse)
	err := c.cc.Invoke(ctx, Query_PriceAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PriceFeeder(ctx context.Context, in *QueryGetPriceFeederRequest, opts ...grpc.CallOption) (*QueryGetPriceFeederResponse, error) {
	out := new(QueryGetPriceFeederResponse)
	err := c.cc.Invoke(ctx, Query_PriceFeeder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PriceFeederAll(ctx context.Context, in *QueryAllPriceFeederRequest, opts ...grpc.CallOption) (*QueryAllPriceFeederResponse, error) {
	out := new(QueryAllPriceFeederResponse)
	err := c.cc.Invoke(ctx, Query_PriceFeederAll_FullMethodName, in, out, opts...)
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
	// BandPriceResult defines a rpc handler method for MsgRequestBandPrice.
	BandPriceResult(context.Context, *QueryBandPriceRequest) (*QueryBandPriceResponse, error)
	// LastBandRequestId query the last BandPrice result id
	LastBandRequestId(context.Context, *QueryLastBandRequestIdRequest) (*QueryLastBandRequestIdResponse, error)
	// Queries a AssetInfo by denom.
	AssetInfo(context.Context, *QueryGetAssetInfoRequest) (*QueryGetAssetInfoResponse, error)
	// Queries a list of AssetInfo items.
	AssetInfoAll(context.Context, *QueryAllAssetInfoRequest) (*QueryAllAssetInfoResponse, error)
	// Queries a Price by asset.
	Price(context.Context, *QueryGetPriceRequest) (*QueryGetPriceResponse, error)
	// Queries a list of Price items.
	PriceAll(context.Context, *QueryAllPriceRequest) (*QueryAllPriceResponse, error)
	// Queries a PriceFeeder by feeder.
	PriceFeeder(context.Context, *QueryGetPriceFeederRequest) (*QueryGetPriceFeederResponse, error)
	// Queries a list of PriceFeeder items.
	PriceFeederAll(context.Context, *QueryAllPriceFeederRequest) (*QueryAllPriceFeederResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) BandPriceResult(context.Context, *QueryBandPriceRequest) (*QueryBandPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BandPriceResult not implemented")
}
func (UnimplementedQueryServer) LastBandRequestId(context.Context, *QueryLastBandRequestIdRequest) (*QueryLastBandRequestIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastBandRequestId not implemented")
}
func (UnimplementedQueryServer) AssetInfo(context.Context, *QueryGetAssetInfoRequest) (*QueryGetAssetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetInfo not implemented")
}
func (UnimplementedQueryServer) AssetInfoAll(context.Context, *QueryAllAssetInfoRequest) (*QueryAllAssetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetInfoAll not implemented")
}
func (UnimplementedQueryServer) Price(context.Context, *QueryGetPriceRequest) (*QueryGetPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Price not implemented")
}
func (UnimplementedQueryServer) PriceAll(context.Context, *QueryAllPriceRequest) (*QueryAllPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PriceAll not implemented")
}
func (UnimplementedQueryServer) PriceFeeder(context.Context, *QueryGetPriceFeederRequest) (*QueryGetPriceFeederResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PriceFeeder not implemented")
}
func (UnimplementedQueryServer) PriceFeederAll(context.Context, *QueryAllPriceFeederRequest) (*QueryAllPriceFeederResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PriceFeederAll not implemented")
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

func _Query_BandPriceResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBandPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BandPriceResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BandPriceResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BandPriceResult(ctx, req.(*QueryBandPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LastBandRequestId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLastBandRequestIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LastBandRequestId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_LastBandRequestId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LastBandRequestId(ctx, req.(*QueryLastBandRequestIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AssetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAssetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AssetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AssetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AssetInfo(ctx, req.(*QueryGetAssetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AssetInfoAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllAssetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AssetInfoAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AssetInfoAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AssetInfoAll(ctx, req.(*QueryAllAssetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Price_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Price(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Price_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Price(ctx, req.(*QueryGetPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PriceAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PriceAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PriceAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PriceAll(ctx, req.(*QueryAllPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PriceFeeder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPriceFeederRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PriceFeeder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PriceFeeder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PriceFeeder(ctx, req.(*QueryGetPriceFeederRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PriceFeederAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPriceFeederRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PriceFeederAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PriceFeederAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PriceFeederAll(ctx, req.(*QueryAllPriceFeederRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "elys.oracle.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "BandPriceResult",
			Handler:    _Query_BandPriceResult_Handler,
		},
		{
			MethodName: "LastBandRequestId",
			Handler:    _Query_LastBandRequestId_Handler,
		},
		{
			MethodName: "AssetInfo",
			Handler:    _Query_AssetInfo_Handler,
		},
		{
			MethodName: "AssetInfoAll",
			Handler:    _Query_AssetInfoAll_Handler,
		},
		{
			MethodName: "Price",
			Handler:    _Query_Price_Handler,
		},
		{
			MethodName: "PriceAll",
			Handler:    _Query_PriceAll_Handler,
		},
		{
			MethodName: "PriceFeeder",
			Handler:    _Query_PriceFeeder_Handler,
		},
		{
			MethodName: "PriceFeederAll",
			Handler:    _Query_PriceFeederAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elys/oracle/query.proto",
}
