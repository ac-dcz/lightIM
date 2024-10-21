package mq

import (
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/context"
	"lightIM/edge/tcpedge/internal/svc"
)

type Mq struct {
	svcCtx  *svc.ServiceContext
	handler *MqHandler
}

func NewMq(svcCtx *svc.ServiceContext) *Mq {
	return &Mq{
		svcCtx:  svcCtx,
		handler: defaultHandler,
	}
}

func NewMqWithHandler(svcCtx *svc.ServiceContext, handler *MqHandler) *Mq {
	return &Mq{
		svcCtx:  svcCtx,
		handler: handler,
	}
}

func (m *Mq) Start() {
	defer m.svcCtx.ChatMq.Close()
	for {
		msg, err := m.svcCtx.ChatMq.Fetch(context.TODO())
		if err != nil {
			logx.Errorf("Mq fetch message error: %v", err)
			return
		}
		if err := m.handler.HandleMessage(&msg, func(msg *kafka.Message, err error) {
			if err == nil {
				_ = m.svcCtx.ChatMq.Commit(context.TODO(), *msg)
			} else {
				//TODO: try again
			}
		}); err != nil {
			logx.Errorf("Mq handle message error: %v", err)
			return
		}
	}
}

func (m *Mq) Stop() error {
	return m.svcCtx.ChatMq.Close()
}
