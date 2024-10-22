package logic

import (
	"context"

	"lightIM/rpc/message/internal/svc"
	"lightIM/rpc/message/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupHistoryLogic {
	return &GetGroupHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupHistoryLogic) GetGroupHistory(in *types.GroupHistoryReq) (*types.GroupHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &types.GroupHistoryResp{}, nil
}
