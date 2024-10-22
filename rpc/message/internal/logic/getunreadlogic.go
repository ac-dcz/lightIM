package logic

import (
	"context"

	"lightIM/rpc/message/internal/svc"
	"lightIM/rpc/message/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUnReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnReadLogic {
	return &GetUnReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUnReadLogic) GetUnRead(in *types.UnReadReq) (*types.UnReadResp, error) {
	// todo: add your logic here and delete this line

	return &types.UnReadResp{}, nil
}
