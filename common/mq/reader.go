package mq

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

type ReaderConf struct {
	Brokers   []string `json:"brokers"`
	Topic     string   `json:"topic"`
	GroupName string   `json:"group_name"`
}

type Reader struct {
	r *kafka.Reader
}

func NewReader(conf *ReaderConf) *Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        conf.Brokers,
		Topic:          conf.Topic,
		GroupID:        conf.GroupName,
		CommitInterval: time.Second,
		Logger:         log.New(os.Stdout, "[kafka]", log.LstdFlags),
		ErrorLogger:    log.New(os.Stdout, "[kafka]", log.LstdFlags),
	})
	return &Reader{r: r}
}

func (r *Reader) Read(ctx context.Context) (kafka.Message, error) {
	for {
		msg, err := r.r.ReadMessage(ctx)
		if errors.Is(err, kafka.LeaderNotAvailable) {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		if err != nil {
			return kafka.Message{}, err
		}
		return msg, nil
	}
}

func (r *Reader) Fetch(ctx context.Context) (kafka.Message, error) {
	for {
		msg, err := r.r.FetchMessage(ctx)
		if errors.Is(err, kafka.LeaderNotAvailable) {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		if err != nil {
			return kafka.Message{}, err
		}
		return msg, nil
	}
}

func (r *Reader) Commit(ctx context.Context, msg ...kafka.Message) error {
	return r.r.CommitMessages(ctx, msg...)
}

func (r *Reader) Close() error {
	return r.r.Close()
}
