package logic

import (
	"context"
	"lightIM/common/codes"
	"lightIM/common/params"

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
	//Step1: Update Message State

	for _, id := range in.MsgId {
		if err := l.svcCtx.MessageModel.UpdateStatus(l.ctx, id, params.Read); err != nil {
			l.Logger.Errorf("MongoDB Update Status error: %v", err)
			return &types.AckResp{
				Base: &types.Base{
					Code: codes.InternalServerError.Code,
					Msg:  err.Error(),
				},
			}, nil
		}
	}
	//Step2: Remove UnRead
	if err := l.svcCtx.HistoryModel.RemoveUnRead(l.ctx, in.Uid, in.MsgId...); err != nil {
		l.Logger.Errorf("MongoDB Remove UnRead Message error: %v", err)
		return &types.AckResp{
			Base: &types.Base{
				Code: codes.InternalServerError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}

	return &types.AckResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}
