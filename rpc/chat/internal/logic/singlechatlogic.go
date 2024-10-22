package logic

import (
	"context"

	"lightIM/rpc/chat/internal/svc"
	"lightIM/rpc/chat/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SingleChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSingleChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SingleChatLogic {
	return &SingleChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SingleChatLogic) SingleChat(in *types.SingleChatReq) (*types.SingleChatResp, error) {
	// todo: add your logic here and delete this line
	// Step1: rpc message save in mongodb

	// Step2: push kafka notify edge
	return &types.SingleChatResp{}, nil
}
