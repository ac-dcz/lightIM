package logic

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/mq"
	"lightIM/common/params"
	"lightIM/common/sd"
	"lightIM/rpc/chat/internal/svc"
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

func readerConfkey(readConf mq.ReaderConf) string {
	return fmt.Sprintf("%s#%s", strings.Join(readConf.Brokers, "-"), readConf.Topic)
}

func (l *ProduceLogic) Produce(ctx context.Context, metaData sd.MetaData, msg kafka.Message) error {
	readConf := metaData[params.EdgeTcpServer.EtcdEdgeKq].(mq.ReaderConf)
	if err := l.svcCtx.Producer.Produce(ctx, readerConfkey(readConf), mq.WriterConf{
		Brokers: readConf.Brokers,
		Topic:   readerConfkey(readConf),
	}, msg); err != nil {
		l.Logger.Errorf("Produce error: %v", err)
		return err
	}
	return nil
}
