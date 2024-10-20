package logic

import (
	"context"

	"lightIM/rpc/online/internal/svc"
	"lightIM/rpc/online/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOnlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOnlineLogic {
	return &UserOnlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOnlineLogic) UserOnline(in *types.UserOfflineReq) (*types.UserOfflineResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserOfflineResp{}, nil
}
