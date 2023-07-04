// Code generated by goctl. DO NOT EDIT.
// Source: reply.proto

package reply

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Reply interface {
		Comments(ctx context.Context, in *CommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error)
	}

	defaultReply struct {
		cli zrpc.Client
	}
)

func NewReply(cli zrpc.Client) Reply {
	return &defaultReply{
		cli: cli,
	}
}

func (m *defaultReply) Comments(ctx context.Context, in *CommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error) {
	client := NewReplyClient(m.cli.Conn())
	return client.Comments(ctx, in, opts...)
}
