package logic

import (
	"context"

	"lightIM/rpc/message/internal/svc"
	"lightIM/rpc/message/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AckMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAckMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AckMsgLogic {
	return &AckMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AckMsgLogic) AckMsg(in *types.AckReq) (*types.AckResp, error) {
	// todo: add your logic here and delete this line

	return &types.AckResp{}, nil
}
