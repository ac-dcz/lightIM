package logic

import (
	"context"

	"lightIM/rpc/chat/internal/svc"
	"lightIM/rpc/chat/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupChatLogic {
	return &GroupChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupChatLogic) GroupChat(in *types.GroupChatReq) (*types.GroupChatResp, error) {
	// todo: add your logic here and delete this line

	return &types.GroupChatResp{}, nil
}
