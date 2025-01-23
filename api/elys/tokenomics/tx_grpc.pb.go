// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: elys/tokenomics/tx.proto

package tokenomics

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
	Msg_CreateAirdrop_FullMethodName            = "/elys.tokenomics.Msg/CreateAirdrop"
	Msg_UpdateAirdrop_FullMethodName            = "/elys.tokenomics.Msg/UpdateAirdrop"
	Msg_DeleteAirdrop_FullMethodName            = "/elys.tokenomics.Msg/DeleteAirdrop"
	Msg_ClaimAirdrop_FullMethodName             = "/elys.tokenomics.Msg/ClaimAirdrop"
	Msg_UpdateGenesisInflation_FullMethodName   = "/elys.tokenomics.Msg/UpdateGenesisInflation"
	Msg_CreateTimeBasedInflation_FullMethodName = "/elys.tokenomics.Msg/CreateTimeBasedInflation"
	Msg_UpdateTimeBasedInflation_FullMethodName = "/elys.tokenomics.Msg/UpdateTimeBasedInflation"
	Msg_DeleteTimeBasedInflation_FullMethodName = "/elys.tokenomics.Msg/DeleteTimeBasedInflation"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	CreateAirdrop(ctx context.Context, in *MsgCreateAirdrop, opts ...grpc.CallOption) (*MsgCreateAirdropResponse, error)
	UpdateAirdrop(ctx context.Context, in *MsgUpdateAirdrop, opts ...grpc.CallOption) (*MsgUpdateAirdropResponse, error)
	DeleteAirdrop(ctx context.Context, in *MsgDeleteAirdrop, opts ...grpc.CallOption) (*MsgDeleteAirdropResponse, error)
	ClaimAirdrop(ctx context.Context, in *MsgClaimAirdrop, opts ...grpc.CallOption) (*MsgClaimAirdropResponse, error)
	UpdateGenesisInflation(ctx context.Context, in *MsgUpdateGenesisInflation, opts ...grpc.CallOption) (*MsgUpdateGenesisInflationResponse, error)
	CreateTimeBasedInflation(ctx context.Context, in *MsgCreateTimeBasedInflation, opts ...grpc.CallOption) (*MsgCreateTimeBasedInflationResponse, error)
	UpdateTimeBasedInflation(ctx context.Context, in *MsgUpdateTimeBasedInflation, opts ...grpc.CallOption) (*MsgUpdateTimeBasedInflationResponse, error)
	DeleteTimeBasedInflation(ctx context.Context, in *MsgDeleteTimeBasedInflation, opts ...grpc.CallOption) (*MsgDeleteTimeBasedInflationResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateAirdrop(ctx context.Context, in *MsgCreateAirdrop, opts ...grpc.CallOption) (*MsgCreateAirdropResponse, error) {
	out := new(MsgCreateAirdropResponse)
	err := c.cc.Invoke(ctx, Msg_CreateAirdrop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateAirdrop(ctx context.Context, in *MsgUpdateAirdrop, opts ...grpc.CallOption) (*MsgUpdateAirdropResponse, error) {
	out := new(MsgUpdateAirdropResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateAirdrop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteAirdrop(ctx context.Context, in *MsgDeleteAirdrop, opts ...grpc.CallOption) (*MsgDeleteAirdropResponse, error) {
	out := new(MsgDeleteAirdropResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteAirdrop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimAirdrop(ctx context.Context, in *MsgClaimAirdrop, opts ...grpc.CallOption) (*MsgClaimAirdropResponse, error) {
	out := new(MsgClaimAirdropResponse)
	err := c.cc.Invoke(ctx, Msg_ClaimAirdrop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGenesisInflation(ctx context.Context, in *MsgUpdateGenesisInflation, opts ...grpc.CallOption) (*MsgUpdateGenesisInflationResponse, error) {
	out := new(MsgUpdateGenesisInflationResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateGenesisInflation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateTimeBasedInflation(ctx context.Context, in *MsgCreateTimeBasedInflation, opts ...grpc.CallOption) (*MsgCreateTimeBasedInflationResponse, error) {
	out := new(MsgCreateTimeBasedInflationResponse)
	err := c.cc.Invoke(ctx, Msg_CreateTimeBasedInflation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateTimeBasedInflation(ctx context.Context, in *MsgUpdateTimeBasedInflation, opts ...grpc.CallOption) (*MsgUpdateTimeBasedInflationResponse, error) {
	out := new(MsgUpdateTimeBasedInflationResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateTimeBasedInflation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteTimeBasedInflation(ctx context.Context, in *MsgDeleteTimeBasedInflation, opts ...grpc.CallOption) (*MsgDeleteTimeBasedInflationResponse, error) {
	out := new(MsgDeleteTimeBasedInflationResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteTimeBasedInflation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	CreateAirdrop(context.Context, *MsgCreateAirdrop) (*MsgCreateAirdropResponse, error)
	UpdateAirdrop(context.Context, *MsgUpdateAirdrop) (*MsgUpdateAirdropResponse, error)
	DeleteAirdrop(context.Context, *MsgDeleteAirdrop) (*MsgDeleteAirdropResponse, error)
	ClaimAirdrop(context.Context, *MsgClaimAirdrop) (*MsgClaimAirdropResponse, error)
	UpdateGenesisInflation(context.Context, *MsgUpdateGenesisInflation) (*MsgUpdateGenesisInflationResponse, error)
	CreateTimeBasedInflation(context.Context, *MsgCreateTimeBasedInflation) (*MsgCreateTimeBasedInflationResponse, error)
	UpdateTimeBasedInflation(context.Context, *MsgUpdateTimeBasedInflation) (*MsgUpdateTimeBasedInflationResponse, error)
	DeleteTimeBasedInflation(context.Context, *MsgDeleteTimeBasedInflation) (*MsgDeleteTimeBasedInflationResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateAirdrop(context.Context, *MsgCreateAirdrop) (*MsgCreateAirdropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAirdrop not implemented")
}
func (UnimplementedMsgServer) UpdateAirdrop(context.Context, *MsgUpdateAirdrop) (*MsgUpdateAirdropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAirdrop not implemented")
}
func (UnimplementedMsgServer) DeleteAirdrop(context.Context, *MsgDeleteAirdrop) (*MsgDeleteAirdropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAirdrop not implemented")
}
func (UnimplementedMsgServer) ClaimAirdrop(context.Context, *MsgClaimAirdrop) (*MsgClaimAirdropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimAirdrop not implemented")
}
func (UnimplementedMsgServer) UpdateGenesisInflation(context.Context, *MsgUpdateGenesisInflation) (*MsgUpdateGenesisInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGenesisInflation not implemented")
}
func (UnimplementedMsgServer) CreateTimeBasedInflation(context.Context, *MsgCreateTimeBasedInflation) (*MsgCreateTimeBasedInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTimeBasedInflation not implemented")
}
func (UnimplementedMsgServer) UpdateTimeBasedInflation(context.Context, *MsgUpdateTimeBasedInflation) (*MsgUpdateTimeBasedInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTimeBasedInflation not implemented")
}
func (UnimplementedMsgServer) DeleteTimeBasedInflation(context.Context, *MsgDeleteTimeBasedInflation) (*MsgDeleteTimeBasedInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTimeBasedInflation not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateAirdrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAirdrop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAirdrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateAirdrop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAirdrop(ctx, req.(*MsgCreateAirdrop))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateAirdrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateAirdrop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateAirdrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateAirdrop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateAirdrop(ctx, req.(*MsgUpdateAirdrop))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteAirdrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteAirdrop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteAirdrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteAirdrop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteAirdrop(ctx, req.(*MsgDeleteAirdrop))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimAirdrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimAirdrop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimAirdrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ClaimAirdrop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimAirdrop(ctx, req.(*MsgClaimAirdrop))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGenesisInflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGenesisInflation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGenesisInflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateGenesisInflation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGenesisInflation(ctx, req.(*MsgUpdateGenesisInflation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateTimeBasedInflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateTimeBasedInflation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateTimeBasedInflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateTimeBasedInflation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateTimeBasedInflation(ctx, req.(*MsgCreateTimeBasedInflation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateTimeBasedInflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateTimeBasedInflation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateTimeBasedInflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateTimeBasedInflation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateTimeBasedInflation(ctx, req.(*MsgUpdateTimeBasedInflation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteTimeBasedInflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteTimeBasedInflation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteTimeBasedInflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteTimeBasedInflation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteTimeBasedInflation(ctx, req.(*MsgDeleteTimeBasedInflation))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "elys.tokenomics.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAirdrop",
			Handler:    _Msg_CreateAirdrop_Handler,
		},
		{
			MethodName: "UpdateAirdrop",
			Handler:    _Msg_UpdateAirdrop_Handler,
		},
		{
			MethodName: "DeleteAirdrop",
			Handler:    _Msg_DeleteAirdrop_Handler,
		},
		{
			MethodName: "ClaimAirdrop",
			Handler:    _Msg_ClaimAirdrop_Handler,
		},
		{
			MethodName: "UpdateGenesisInflation",
			Handler:    _Msg_UpdateGenesisInflation_Handler,
		},
		{
			MethodName: "CreateTimeBasedInflation",
			Handler:    _Msg_CreateTimeBasedInflation_Handler,
		},
		{
			MethodName: "UpdateTimeBasedInflation",
			Handler:    _Msg_UpdateTimeBasedInflation_Handler,
		},
		{
			MethodName: "DeleteTimeBasedInflation",
			Handler:    _Msg_DeleteTimeBasedInflation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elys/tokenomics/tx.proto",
}
