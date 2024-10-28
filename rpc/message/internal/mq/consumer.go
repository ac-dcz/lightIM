package mq

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/mq"
	"lightIM/rpc/message/internal/svc"
)

type ConsumerOptions struct {
	PoolSize int
}

type Consumer struct {
	svcCtx  *svc.ServiceContext
	handler *Handler
	reader  *mq.Reader
}

func NewConsumer(svcCtx *svc.ServiceContext, opt *ConsumerOptions) (*Consumer, error) {
	reader := mq.NewReader(&svcCtx.Config.MsgReader)
	handler, err := NewHandler(svcCtx, opt)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		reader:  reader,
		handler: handler,
		svcCtx:  svcCtx,
	}, nil
}

func (c *Consumer) Start() {
	defer c.reader.Close()
	for {
		msg, err := c.reader.Fetch(context.Background())
		if err != nil {
			logx.Errorf("ImConsumer fetch message error: %v", err)
			return
		}
		if err = c.handler.HandleMessage(&msg, func(msg *kafka.Message, err error) {
			if err == nil {
				_ = c.reader.Commit(context.TODO(), *msg)
			} else {
				//TODO: try again
			}
		}); err != nil {
			logx.Errorf("ImConsumer handle message error: %v", err)
			return
		}
	}
}

func (c *Consumer) Stop() error {
	defer c.reader.Close()
	if c.handler != nil {
		return c.handler.Close()
	}
	return nil
}
