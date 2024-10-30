package logic

import (
	"context"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupAckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAckLogic {
	return &GroupAckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupAckLogic) GroupAck(in *types.JoinGroupAck) (*types.JoinGroupAckResp, error) {
	// todo: add your logic here and delete this line

	return &types.JoinGroupAckResp{}, nil
}
