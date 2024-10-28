package mq

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

type WriterConf struct {
	Brokers []string `json:"brokers"`
	Topic   string   `json:"topic"`
}

type Writer struct {
	w *kafka.Writer
}

func NewWriterSync(conf *WriterConf) *Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(conf.Brokers...),
		Topic:                  conf.Topic,
		AllowAutoTopicCreation: true,
		RequiredAcks:           kafka.RequireOne,
		Balancer:               &kafka.Hash{},
		Logger:                 log.New(os.Stdout, "[kafka]", log.LstdFlags),
		ErrorLogger:            log.New(os.Stdout, "[kafka]", log.LstdFlags),
	}
	return &Writer{w: w}
}

func NewWriterAsync(conf *WriterConf, Completion func(messages []kafka.Message, err error)) *Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(conf.Brokers...),
		Topic:                  conf.Topic,
		AllowAutoTopicCreation: true,
		RequiredAcks:           kafka.RequireOne,
		Async:                  true,
		Completion:             Completion,
		Logger:                 log.New(os.Stdout, "[kafka]", log.LstdFlags),
		ErrorLogger:            log.New(os.Stdout, "[kafka]", log.LstdFlags),
	}
	return &Writer{w: w}
}

func (w *Writer) Write(ctx context.Context, msg ...kafka.Message) error {
	return w.w.WriteMessages(ctx, msg...)
}

func (w *Writer) Close() error {
	return w.Close()
}
