// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: raft.proto

package raft

import (
	context "context"
	raftpb "github.com/vanus-labs/vanus/raft/raftpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RaftServerClient is the client API for RaftServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftServerClient interface {
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (RaftServer_SendMessageClient, error)
}

type raftServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftServerClient(cc grpc.ClientConnInterface) RaftServerClient {
	return &raftServerClient{cc}
}

func (c *raftServerClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (RaftServer_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &RaftServer_ServiceDesc.Streams[0], "/vanus.core.raft.RaftServer/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftServerSendMessageClient{stream}
	return x, nil
}

type RaftServer_SendMessageClient interface {
	Send(*raftpb.Message) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type raftServerSendMessageClient struct {
	grpc.ClientStream
}

func (x *raftServerSendMessageClient) Send(m *raftpb.Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *raftServerSendMessageClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RaftServerServer is the server API for RaftServer service.
// All implementations should embed UnimplementedRaftServerServer
// for forward compatibility
type RaftServerServer interface {
	SendMessage(RaftServer_SendMessageServer) error
}

// UnimplementedRaftServerServer should be embedded to have forward compatible implementations.
type UnimplementedRaftServerServer struct {
}

func (UnimplementedRaftServerServer) SendMessage(RaftServer_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

// UnsafeRaftServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftServerServer will
// result in compilation errors.
type UnsafeRaftServerServer interface {
	mustEmbedUnimplementedRaftServerServer()
}

func RegisterRaftServerServer(s grpc.ServiceRegistrar, srv RaftServerServer) {
	s.RegisterService(&RaftServer_ServiceDesc, srv)
}

func _RaftServer_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RaftServerServer).SendMessage(&raftServerSendMessageServer{stream})
}

type RaftServer_SendMessageServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*raftpb.Message, error)
	grpc.ServerStream
}

type raftServerSendMessageServer struct {
	grpc.ServerStream
}

func (x *raftServerSendMessageServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *raftServerSendMessageServer) Recv() (*raftpb.Message, error) {
	m := new(raftpb.Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RaftServer_ServiceDesc is the grpc.ServiceDesc for RaftServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vanus.core.raft.RaftServer",
	HandlerType: (*RaftServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _RaftServer_SendMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "raft.proto",
}
