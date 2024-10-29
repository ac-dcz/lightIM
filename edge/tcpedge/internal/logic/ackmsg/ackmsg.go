package ackmsg

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"lightIM/rpc/message/message"
)

type AckMsgLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAckMsgLogic(svcCtx *svc.ServiceContext) *AckMsgLogic {
	return &AckMsgLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *AckMsgLogic) AckMsg(ctx context.Context, msg *types.AckMsg) (*types.AckMsgResp, error) {
	if resp, err := l.svcCtx.MessageRpc.AckMsg(ctx, &message.AckReq{
		Uid:   msg.From,
		MsgId: msg.Ack,
	}); err != nil {
		l.Logger.Errorf("ack msg error: %v", err)
		return nil, err
	} else if resp.Base.Code != codes.OK.Code {
		l.Logger.Errorf("ack msg error: %#v", resp.Base)
		return nil, errors.New(resp.Base.Msg)
	}
	return &types.AckMsgResp{
		RespBase: types.RespBase{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}
