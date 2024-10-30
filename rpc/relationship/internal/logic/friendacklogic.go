package logic

import (
	"context"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendAckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendAckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAckLogic {
	return &FriendAckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendAckLogic) FriendAck(in *types.AddFriendAck) (*types.AddFriendAckResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddFriendAckResp{}, nil
}
