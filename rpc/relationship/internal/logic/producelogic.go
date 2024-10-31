package logic

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/mq"
	"lightIM/rpc/relationship/internal/svc"
	"strings"
)

type ProduceLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProduceLogic(svcCtx *svc.ServiceContext) *ProduceLogic {
	return &ProduceLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func readerConfkey(readConf *mq.ReaderConf) string {
	return fmt.Sprintf("%s#%s", strings.Join(readConf.Brokers, "-"), readConf.Topic)
}

func (l *ProduceLogic) Produce(ctx context.Context, rdConf *mq.ReaderConf, msg kafka.Message) error {
	if err := l.svcCtx.Producer.Produce(ctx, readerConfkey(rdConf), mq.WriterConf{
		Brokers: rdConf.Brokers,
		Topic:   rdConf.Topic,
	}, msg); err != nil {
		l.Logger.Errorf("Produce error: %v", err)
		return err
	}
	return nil
}
