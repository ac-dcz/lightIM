package mq

import (
	"context"
	"github.com/segmentio/kafka-go"
	"lightIM/common/mq"
	"lightIM/rpc/relationship/internal/svc"
	"sync"
)

type ConsumerOptions struct {
	PoolSize int
}

type Consumer struct {
	svcCtx         *svc.ServiceContext
	handler        mq.Handler
	onlineConsumer *mq.Consumer
}

func NewConsumer(svcCtx *svc.ServiceContext, opt *ConsumerOptions) (*Consumer, error) {
	handler, err := NewHandler(svcCtx, opt)
	if err != nil {
		return nil, err
	}
	onlineReader := mq.NewReader(&svcCtx.Config.KqOnlineReader)
	return &Consumer{
		onlineConsumer: mq.NewConsumer(handler, onlineReader, func(msg *kafka.Message, err error) {
			if err == nil {
				_ = onlineReader.Commit(context.TODO(), *msg)
			} else {
				//TODO: try again
			}
		}),
		handler: handler,
		svcCtx:  svcCtx,
	}, nil
}

func (c *Consumer) Start() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = c.onlineConsumer.Consume()
	}()
	wg.Wait()
}

func (c *Consumer) Stop() error {
	_ = c.onlineConsumer.Close()
	return nil
}
