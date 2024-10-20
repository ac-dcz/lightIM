package logic

import (
	"context"

	"lightIM/rpc/online/internal/svc"
	"lightIM/rpc/online/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRouteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRouteLogic {
	return &GetRouteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRouteLogic) GetRoute(in *types.RouteReq) (*types.RouteResp, error) {
	// todo: add your logic here and delete this line

	return &types.RouteResp{}, nil
}
