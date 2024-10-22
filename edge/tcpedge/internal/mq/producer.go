package mq

import (
	"context"
	"github.com/segmentio/kafka-go"
	"lightIM/common/mq"
	"lightIM/rpc/message/mqtypes"
)

type ImProducer struct {
	writerMq *mq.Writer
}

func NewImProducer(writerMq *mq.Writer) *ImProducer {
	return &ImProducer{
		writerMq: writerMq,
	}
}

func (p *ImProducer) Write(ctx context.Context, msg *mqtypes.Message) error {
	if data, err := msg.Encode(); err != nil {
		return err
	} else {
		if err = p.writerMq.Write(ctx, kafka.Message{
			Value: data,
		}); err != nil {
			return err
		}
	}
	return nil
}
