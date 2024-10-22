package singlechat

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

func (l *SingleChatLogic) SingleChat(msg *types.SingleChatMsg, _ *imnet.ImConn) (*types.SingleChatMsgResp, error) {
	// 文本消息
	if msg.Content.Type == params.Text {
		text := msg.Content.Data.(*types.TextContent)

		msgMq := &mqtypes.Message{
			Type:         msg.Content.Type,
			Status:       params.UnRead,
			From:         msg.From,
			To:           msg.To,
			IsGroup:      false,
			EncodingType: text.EncodingType,
			Data:         text.Body,
			TimeStamp:    time.Now().Unix(),
		}

		if err := l.svcCtx.ChatProducer.Write(context.Background(), msgMq); err != nil {
			l.Logger.Errorf("write to kafka error: %v", err)
			return &types.SingleChatMsgResp{
				RespBase: types.RespBase{
					Code: codes.InternalServerError.Code,
					Msg:  codes.InternalServerError.Msg,
				},
			}, nil
		}

		return &types.SingleChatMsgResp{
			RespBase: types.RespBase{
				Code: codes.OK.Code,
				Msg:  codes.OK.Msg,
			},
		}, nil
	}
	return &types.SingleChatMsgResp{}, nil
}
