package logic

import (
	"context"

	"lightIM/rpc/message/internal/svc"
	"lightIM/rpc/message/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMsgStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMsgStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMsgStatusLogic {
	return &UpdateMsgStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMsgStatusLogic) UpdateMsgStatus(in *types.UpdateMsgStatusReq) (*types.UpdateMsgStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateMsgStatusResp{}, nil
}
