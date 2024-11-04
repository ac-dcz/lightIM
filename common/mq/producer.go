package mq

import (
	"context"
	"github.com/segmentio/kafka-go"
	"sync"
)

type Producer struct {
	mu      sync.Mutex
	writers map[string]*Writer
}

func NewProducer() *Producer {
	return &Producer{writers: make(map[string]*Writer)}
}

func (p *Producer) writer(key string, conf WriterConf) *Writer {
	p.mu.Lock()
	defer p.mu.Unlock()
	if w, ok := p.writers[key]; ok {
		return w
	} else {
		p.writers[key] = NewWriterSync(&conf)
	}
	return p.writers[key]
}

func (p *Producer) Produce(ctx context.Context, key string, conf WriterConf, msg kafka.Message) error {
	w := p.writer(key, conf)
	return w.Write(ctx, msg)
}
