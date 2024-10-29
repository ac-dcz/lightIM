package consumer

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/params"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"lightIM/rpc/message/mqtypes"
)

type Logic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogic(svcCtx *svc.ServiceContext) *Logic {
	return &Logic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *Logic) Exec(msg *mqtypes.Message) error {
	switch msg.Type {
	case params.Text:
		return l.execTextMsg(msg)
	case params.NormalFile:
	case params.BigFile:
	}
	return nil
}

func (l *Logic) execTextMsg(msg *mqtypes.Message) error {
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

func (l *Logic) execBigFileMsg(msg *mqtypes.Message) error {

	return nil
}

func (l *Logic) execNormalFileMsg(msg *mqtypes.Message) error {

	return nil
}
