// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package server

import (
	"context"

	"lightIM/rpc/chat/internal/logic"
	"lightIM/rpc/chat/internal/svc"
	"lightIM/rpc/chat/types"
)

type ChatServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedChatServer
}

func NewChatServer(svcCtx *svc.ServiceContext) *ChatServer {
	return &ChatServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServer) SingleChat(ctx context.Context, in *types.SingleChatReq) (*types.SingleChatResp, error) {
	l := logic.NewSingleChatLogic(ctx, s.svcCtx)
	return l.SingleChat(in)
}

func (s *ChatServer) GroupChat(ctx context.Context, in *types.GroupChatReq) (*types.GroupChatResp, error) {
	l := logic.NewGroupChatLogic(ctx, s.svcCtx)
	return l.GroupChat(in)
}