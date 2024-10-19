package multichat

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
)

type MultiChatLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMultiChatLogic(svcCtx *svc.ServiceContext) *MultiChatLogic {
	return &MultiChatLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *MultiChatLogic) MultiChat(msg *types.MultiChatMsg) (*types.MultiChatMsgResp, error) {
	return &types.MultiChatMsgResp{}, nil
}
