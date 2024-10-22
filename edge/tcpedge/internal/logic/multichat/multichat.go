package multichat

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/common/params"
	"lightIM/edge/tcpedge/internal/imnet"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"lightIM/rpc/message/mqtypes"
	"time"
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

func (l *MultiChatLogic) MultiChat(msg *types.MultiChatMsg, _ *imnet.ImConn) (*types.MultiChatMsgResp, error) {
	// 文本消息
	if msg.Content.Type == params.Text {
		text := msg.Content.Data.(*types.TextContent)

		msgMq := &mqtypes.Message{
			Type:         msg.Content.Type,
			Status:       params.UnRead,
			From:         msg.From,
			To:           msg.Group,
			IsGroup:      true,
			EncodingType: text.EncodingType,
			Data:         text.Body,
			TimeStamp:    time.Now().Unix(),
		}

		if err := l.svcCtx.ChatProducer.Write(context.Background(), msgMq); err != nil {
			l.Logger.Errorf("write to kafka error: %v", err)
			return &types.MultiChatMsgResp{
				RespBase: types.RespBase{
					Code: codes.InternalServerError.Code,
					Msg:  codes.InternalServerError.Msg,
				},
			}, nil
		}

		return &types.MultiChatMsgResp{
			RespBase: types.RespBase{
				Code: codes.OK.Code,
				Msg:  codes.OK.Msg,
			},
		}, nil
	}
	return &types.MultiChatMsgResp{}, nil
}
