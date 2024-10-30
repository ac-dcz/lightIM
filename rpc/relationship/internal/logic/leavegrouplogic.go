package logic

import (
	"context"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LeaveGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLeaveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveGroupLogic {
	return &LeaveGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LeaveGroupLogic) LeaveGroup(in *types.LeaveGroupReq) (*types.LeaveGroupReq, error) {
	// todo: add your logic here and delete this line

	return &types.LeaveGroupReq{}, nil
}
