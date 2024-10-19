package singlechat

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
)

type SingleChatLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSingleChatLogic(svcCtx *svc.ServiceContext) *SingleChatLogic {
	return &SingleChatLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *SingleChatLogic) SingleChat(msg *types.SingleChatMsg) (*types.SingleChatMsgResp, error) {
	return &types.SingleChatMsgResp{}, nil
}
