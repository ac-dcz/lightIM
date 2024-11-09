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
	if err := l.svcCtx.GroupMemberModel.Delete(l.ctx, in.GId, in.From); err != nil {
		l.Logger.Errorf("")
		return nil, err
	}

	return &types.LeaveGroupReq{}, nil
}
