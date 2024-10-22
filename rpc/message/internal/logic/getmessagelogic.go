package logic

import (
	"context"

	"lightIM/rpc/message/internal/svc"
	"lightIM/rpc/message/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageLogic) GetMessage(in *types.MsgReq) (*types.MsgResp, error) {
	// todo: add your logic here and delete this line

	return &types.MsgResp{}, nil
}
