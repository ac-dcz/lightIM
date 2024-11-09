package consumer

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/params"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"lightIM/rpc/message/mqtypes"
)

type MessageLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogic(svcCtx *svc.ServiceContext) *MessageLogic {
	return &MessageLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *MessageLogic) Exec(msg *mqtypes.Message) error {
	switch msg.Type {
	case params.Text:
		return l.execTextMsg(msg)
	case params.NormalFile:
	case params.BigFile:
	}
	return nil
}

func (l *MessageLogic) execTextMsg(msg *mqtypes.Message) error {
	if conn, ok := l.svcCtx.ConnPool.GetAuthConnByUid(msg.To); ok {
		if msg.IsGroup {
			m := &types.GroupChatMsg{
				Base: types.Base{
					MsgId:     msg.MsgId, //Mongo Id
					TimeStamp: msg.TimeStamp,
				},
				From:  msg.From,
				Group: msg.Group,
				Content: types.Content{
					Type: msg.Type,
					Data: types.TextContent{
						EncodingType: msg.EncodingType,
						Body:         msg.Data,
					},
				},
			}
			if err := conn.Write(m); err != nil {
				l.Logger.Errorf("Conn write group chat msg error: %v", err)
				return err
			}
		} else {
			m := &types.SingleChatMsg{
				Base: types.Base{
					MsgId:     msg.MsgId, //Mongo Id
					TimeStamp: msg.TimeStamp,
				},
				From: msg.From,
				To:   msg.To,
				Content: types.Content{
					Type: msg.Type,
					Data: types.TextContent{
						EncodingType: msg.EncodingType,
						Body:         msg.Data,
					},
				},
			}
			if err := conn.Write(m); err != nil {
				l.Logger.Errorf("Conn write single chat msg error: %v", err)
				return err
			}
		}
	}
	return nil
}

func (l *MessageLogic) execBigFileMsg(msg *mqtypes.Message) error {

	return nil
}

func (l *MessageLogic) execNormalFileMsg(msg *mqtypes.Message) error {

	return nil
}
