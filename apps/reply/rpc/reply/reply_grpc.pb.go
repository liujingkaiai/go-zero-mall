// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: reply.proto

package reply

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
	Reply_Comments_FullMethodName = "/reply.Reply/Comments"
)

// ReplyClient is the client API for Reply service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplyClient interface {
	Comments(ctx context.Context, in *CommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error)
}

type replyClient struct {
	cc grpc.ClientConnInterface
}

func NewReplyClient(cc grpc.ClientConnInterface) ReplyClient {
	return &replyClient{cc}
}

func (c *replyClient) Comments(ctx context.Context, in *CommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error) {
	out := new(CommentsResponse)
	err := c.cc.Invoke(ctx, Reply_Comments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplyServer is the server API for Reply service.
// All implementations must embed UnimplementedReplyServer
// for forward compatibility
type ReplyServer interface {
	Comments(context.Context, *CommentsRequest) (*CommentsResponse, error)
	mustEmbedUnimplementedReplyServer()
}

// UnimplementedReplyServer must be embedded to have forward compatible implementations.
type UnimplementedReplyServer struct {
}

func (UnimplementedReplyServer) Comments(context.Context, *CommentsRequest) (*CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comments not implemented")
}
func (UnimplementedReplyServer) mustEmbedUnimplementedReplyServer() {}

// UnsafeReplyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplyServer will
// result in compilation errors.
type UnsafeReplyServer interface {
	mustEmbedUnimplementedReplyServer()
}

func RegisterReplyServer(s grpc.ServiceRegistrar, srv ReplyServer) {
	s.RegisterService(&Reply_ServiceDesc, srv)
}

func _Reply_Comments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplyServer).Comments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reply_Comments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplyServer).Comments(ctx, req.(*CommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reply_ServiceDesc is the grpc.ServiceDesc for Reply service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reply_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reply.Reply",
	HandlerType: (*ReplyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Comments",
			Handler:    _Reply_Comments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reply.proto",
}
