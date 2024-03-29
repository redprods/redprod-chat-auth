// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: messenger/messenger.proto

package messenger

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

// MessengerServiceClient is the client API for MessengerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerServiceClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	GetConversations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConversationsResponse, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
}

type messengerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerServiceClient(cc grpc.ClientConnInterface) MessengerServiceClient {
	return &messengerServiceClient{cc}
}

func (c *messengerServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/messenger.MessengerService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) GetConversations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConversationsResponse, error) {
	out := new(GetConversationsResponse)
	err := c.cc.Invoke(ctx, "/messenger.MessengerService/GetConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, "/messenger.MessengerService/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServiceServer is the server API for MessengerService service.
// All implementations should embed UnimplementedMessengerServiceServer
// for forward compatibility
type MessengerServiceServer interface {
	SendMessage(context.Context, *Message) (*Message, error)
	GetConversations(context.Context, *emptypb.Empty) (*GetConversationsResponse, error)
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
}

// UnimplementedMessengerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMessengerServiceServer struct {
}

func (UnimplementedMessengerServiceServer) SendMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessengerServiceServer) GetConversations(context.Context, *emptypb.Empty) (*GetConversationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversations not implemented")
}
func (UnimplementedMessengerServiceServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}

// UnsafeMessengerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServiceServer will
// result in compilation errors.
type UnsafeMessengerServiceServer interface {
	mustEmbedUnimplementedMessengerServiceServer()
}

func RegisterMessengerServiceServer(s grpc.ServiceRegistrar, srv MessengerServiceServer) {
	s.RegisterService(&MessengerService_ServiceDesc, srv)
}

func _MessengerService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.MessengerService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_GetConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).GetConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.MessengerService/GetConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).GetConversations(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.MessengerService/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessengerService_ServiceDesc is the grpc.ServiceDesc for MessengerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessengerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.MessengerService",
	HandlerType: (*MessengerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _MessengerService_SendMessage_Handler,
		},
		{
			MethodName: "GetConversations",
			Handler:    _MessengerService_GetConversations_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _MessengerService_GetMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messenger/messenger.proto",
}
