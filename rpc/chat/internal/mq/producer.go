package mq

import (
	"context"
	"github.com/segmentio/kafka-go"
	"lightIM/common/mq"
	"sync"
)

type Producer struct {
	mu      sync.Mutex
	writers map[string]*mq.Writer
}

func NewProducer() *Producer {
	return &Producer{writers: make(map[string]*mq.Writer)}
}

func (p *Producer) writer(key string, conf mq.WriterConf) *mq.Writer {
	if w, ok := p.writers[key]; ok {
		return w
	} else {
		p.writers[key] = mq.NewWriterSync(&conf)
	}
	return p.writers[key]
}

func (p *Producer) Produce(ctx context.Context, key string, conf mq.WriterConf, msg kafka.Message) error {
	w := p.writer(key, conf)
	return w.Write(ctx, msg)
}
