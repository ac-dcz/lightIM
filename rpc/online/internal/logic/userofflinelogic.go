package logic

import (
	"context"

	"lightIM/rpc/online/internal/svc"
	"lightIM/rpc/online/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOfflineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOfflineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOfflineLogic {
	return &UserOfflineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOfflineLogic) UserOffline(in *types.UserOfflineReq) (*types.UserOfflineResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserOfflineResp{}, nil
}
