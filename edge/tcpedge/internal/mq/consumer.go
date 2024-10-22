package mq

import (
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/context"
	"lightIM/common/mq"
	"lightIM/edge/tcpedge/internal/svc"
)

type consumerOptions struct {
	poolSize int
}

type ImConsumer struct {
	svcCtx  *svc.ServiceContext
	handler *consumerHandler
	reader  *mq.Reader
}

func NewImConsumer(svcCtx *svc.ServiceContext, reader *mq.Reader, opt *consumerOptions) (*ImConsumer, error) {
	handler, err := newConsumerHandler(svcCtx, opt)
	if err != nil {
		return nil, err
	}
	return &ImConsumer{
		svcCtx:  svcCtx,
		handler: handler,
		reader:  reader,
	}, nil
}

func (m *ImConsumer) Start() {
	defer m.reader.Close()
	for {
		msg, err := m.reader.Fetch(context.Background())
		if err != nil {
			logx.Errorf("ImConsumer fetch message error: %v", err)
			return
		}
		if err = m.handler.HandleMessage(&msg, func(msg *kafka.Message, err error) {
			if err == nil {
				_ = m.reader.Commit(context.TODO(), *msg)
			} else {
				//TODO: try again
			}
		}); err != nil {
			logx.Errorf("ImConsumer handle message error: %v", err)
			return
		}
	}
}

func (m *ImConsumer) Stop() error {
	return m.reader.Close()
}
