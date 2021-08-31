// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package broadcast

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

// BroadcastServiceClient is the client API for BroadcastService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadcastServiceClient interface {
	SendToAll(ctx context.Context, in *BroadcastMessage, opts ...grpc.CallOption) (*SendToAllResponse, error)
	MessagesFromOthers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BroadcastService_MessagesFromOthersClient, error)
}

type broadcastServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcastServiceClient(cc grpc.ClientConnInterface) BroadcastServiceClient {
	return &broadcastServiceClient{cc}
}

func (c *broadcastServiceClient) SendToAll(ctx context.Context, in *BroadcastMessage, opts ...grpc.CallOption) (*SendToAllResponse, error) {
	out := new(SendToAllResponse)
	err := c.cc.Invoke(ctx, "/broadcast.BroadcastService/SendToAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastServiceClient) MessagesFromOthers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BroadcastService_MessagesFromOthersClient, error) {
	stream, err := c.cc.NewStream(ctx, &BroadcastService_ServiceDesc.Streams[0], "/broadcast.BroadcastService/MessagesFromOthers", opts...)
	if err != nil {
		return nil, err
	}
	x := &broadcastServiceMessagesFromOthersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BroadcastService_MessagesFromOthersClient interface {
	Recv() (*BroadcastMessage, error)
	grpc.ClientStream
}

type broadcastServiceMessagesFromOthersClient struct {
	grpc.ClientStream
}

func (x *broadcastServiceMessagesFromOthersClient) Recv() (*BroadcastMessage, error) {
	m := new(BroadcastMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BroadcastServiceServer is the server API for BroadcastService service.
// All implementations must embed UnimplementedBroadcastServiceServer
// for forward compatibility
type BroadcastServiceServer interface {
	SendToAll(context.Context, *BroadcastMessage) (*SendToAllResponse, error)
	MessagesFromOthers(*Empty, BroadcastService_MessagesFromOthersServer) error
	mustEmbedUnimplementedBroadcastServiceServer()
}

// UnimplementedBroadcastServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBroadcastServiceServer struct {
}

func (UnimplementedBroadcastServiceServer) SendToAll(context.Context, *BroadcastMessage) (*SendToAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToAll not implemented")
}
func (UnimplementedBroadcastServiceServer) MessagesFromOthers(*Empty, BroadcastService_MessagesFromOthersServer) error {
	return status.Errorf(codes.Unimplemented, "method MessagesFromOthers not implemented")
}
func (UnimplementedBroadcastServiceServer) mustEmbedUnimplementedBroadcastServiceServer() {}

// UnsafeBroadcastServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroadcastServiceServer will
// result in compilation errors.
type UnsafeBroadcastServiceServer interface {
	mustEmbedUnimplementedBroadcastServiceServer()
}

func RegisterBroadcastServiceServer(s grpc.ServiceRegistrar, srv BroadcastServiceServer) {
	s.RegisterService(&BroadcastService_ServiceDesc, srv)
}

func _BroadcastService_SendToAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServiceServer).SendToAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/broadcast.BroadcastService/SendToAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServiceServer).SendToAll(ctx, req.(*BroadcastMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BroadcastService_MessagesFromOthers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BroadcastServiceServer).MessagesFromOthers(m, &broadcastServiceMessagesFromOthersServer{stream})
}

type BroadcastService_MessagesFromOthersServer interface {
	Send(*BroadcastMessage) error
	grpc.ServerStream
}

type broadcastServiceMessagesFromOthersServer struct {
	grpc.ServerStream
}

func (x *broadcastServiceMessagesFromOthersServer) Send(m *BroadcastMessage) error {
	return x.ServerStream.SendMsg(m)
}

// BroadcastService_ServiceDesc is the grpc.ServiceDesc for BroadcastService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BroadcastService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "broadcast.BroadcastService",
	HandlerType: (*BroadcastServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendToAll",
			Handler:    _BroadcastService_SendToAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MessagesFromOthers",
			Handler:       _BroadcastService_MessagesFromOthers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol/broadcast.proto",
}
