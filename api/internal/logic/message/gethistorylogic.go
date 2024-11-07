package message

import (
	"context"
	"lightIM/common/codes"
	"lightIM/rpc/message/message"

	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取历史记录
func NewGetHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryLogic {
	return &GetHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHistoryLogic) GetHistory(req *types.HistoryReq) (resp *types.HistoryResp, err error) {
	// todo: add your logic here and delete this line
	if msgResp, e := l.svcCtx.MessageRpc.GetHistory(l.ctx, &message.HistoryReq{
		From: req.Uid1,
		To:   req.Uid2,
	}); e != nil {
		l.Logger.Errorf("GetHistory Rpc err:%v", e)
		return nil, e
	} else if msgResp.Base.Code != codes.OK.Code {
		l.Logger.Errorf("GetHistory Msg resp:%#v", msgResp.Base)
		return &types.HistoryResp{
			Base: types.Base{
				Code: msgResp.Base.Code,
				Msg:  msgResp.Base.Msg,
			},
		}, nil
	} else {
		//var msgList types.Message
	}

	return
}
