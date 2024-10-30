package logic

import (
	"context"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinGroupLogic) JoinGroup(in *types.JoinGroupReq) (*types.JoinGroupResp, error) {
	// todo: add your logic here and delete this line

	return &types.JoinGroupResp{}, nil
}
