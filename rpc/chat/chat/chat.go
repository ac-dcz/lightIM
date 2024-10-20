// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package chat

import (
	"context"

	"lightIM/rpc/chat/types"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Base           = types.Base
	GroupChatReq   = types.GroupChatReq
	GroupChatResp  = types.GroupChatResp
	SingleChatReq  = types.SingleChatReq
	SingleChatResp = types.SingleChatResp

	Chat interface {
		SingleChat(ctx context.Context, in *SingleChatReq, opts ...grpc.CallOption) (*SingleChatResp, error)
		GroupChat(ctx context.Context, in *GroupChatReq, opts ...grpc.CallOption) (*GroupChatResp, error)
	}

	defaultChat struct {
		cli zrpc.Client
	}
)

func NewChat(cli zrpc.Client) Chat {
	return &defaultChat{
		cli: cli,
	}
}

func (m *defaultChat) SingleChat(ctx context.Context, in *SingleChatReq, opts ...grpc.CallOption) (*SingleChatResp, error) {
	client := types.NewChatClient(m.cli.Conn())
	return client.SingleChat(ctx, in, opts...)
}

func (m *defaultChat) GroupChat(ctx context.Context, in *GroupChatReq, opts ...grpc.CallOption) (*GroupChatResp, error) {
	client := types.NewChatClient(m.cli.Conn())
	return client.GroupChat(ctx, in, opts...)
}
