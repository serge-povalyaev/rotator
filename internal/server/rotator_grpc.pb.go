// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: api/rotator.proto

package server

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

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	AddBanner(ctx context.Context, in *AddBannerRequest, opts ...grpc.CallOption) (*Error, error)
	RemoveBanner(ctx context.Context, in *RemoveBannerRequest, opts ...grpc.CallOption) (*Error, error)
	ClickBanner(ctx context.Context, in *ClickBannerRequest, opts ...grpc.CallOption) (*Error, error)
	GetBanner(ctx context.Context, in *GetBannerRequest, opts ...grpc.CallOption) (*GetBannerResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) AddBanner(ctx context.Context, in *AddBannerRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/api.EventService/AddBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) RemoveBanner(ctx context.Context, in *RemoveBannerRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/api.EventService/RemoveBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) ClickBanner(ctx context.Context, in *ClickBannerRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/api.EventService/ClickBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetBanner(ctx context.Context, in *GetBannerRequest, opts ...grpc.CallOption) (*GetBannerResponse, error) {
	out := new(GetBannerResponse)
	err := c.cc.Invoke(ctx, "/api.EventService/GetBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	AddBanner(context.Context, *AddBannerRequest) (*Error, error)
	RemoveBanner(context.Context, *RemoveBannerRequest) (*Error, error)
	ClickBanner(context.Context, *ClickBannerRequest) (*Error, error)
	GetBanner(context.Context, *GetBannerRequest) (*GetBannerResponse, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) AddBanner(context.Context, *AddBannerRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBanner not implemented")
}
func (UnimplementedEventServiceServer) RemoveBanner(context.Context, *RemoveBannerRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBanner not implemented")
}
func (UnimplementedEventServiceServer) ClickBanner(context.Context, *ClickBannerRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClickBanner not implemented")
}
func (UnimplementedEventServiceServer) GetBanner(context.Context, *GetBannerRequest) (*GetBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanner not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_AddBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).AddBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/AddBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).AddBanner(ctx, req.(*AddBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_RemoveBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).RemoveBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/RemoveBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).RemoveBanner(ctx, req.(*RemoveBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_ClickBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClickBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).ClickBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/ClickBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).ClickBanner(ctx, req.(*ClickBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/GetBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetBanner(ctx, req.(*GetBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBanner",
			Handler:    _EventService_AddBanner_Handler,
		},
		{
			MethodName: "RemoveBanner",
			Handler:    _EventService_RemoveBanner_Handler,
		},
		{
			MethodName: "ClickBanner",
			Handler:    _EventService_ClickBanner_Handler,
		},
		{
			MethodName: "GetBanner",
			Handler:    _EventService_GetBanner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/rotator.proto",
}
