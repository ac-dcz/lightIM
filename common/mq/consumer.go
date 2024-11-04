package mq

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type CallBackFunc func(msg *kafka.Message, err error)

type Handler interface {
	HandleMessage(msg *kafka.Message, callBack CallBackFunc)
}

type Consumer struct {
	handler  Handler
	reader   *Reader
	callBack CallBackFunc
}

func NewConsumer(handler Handler, reader *Reader, callback CallBackFunc) *Consumer {
	return &Consumer{
		handler:  handler,
		reader:   reader,
		callBack: callback,
	}
}

func (c *Consumer) Consume() error {
	defer func() {
		_ = c.reader.Close()
	}()

	for {
		msg, err := c.reader.Fetch(context.Background())
		if err != nil {
			return err
		}
		c.handler.HandleMessage(&msg, c.callBack)
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
