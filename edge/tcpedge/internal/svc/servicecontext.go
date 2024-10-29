package svc

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	mq2 "lightIM/common/mq"
	"lightIM/edge/tcpedge/internal/config"
	"lightIM/edge/tcpedge/internal/imnet"
	"lightIM/edge/tcpedge/internal/mq"
	"lightIM/rpc/message/message"
	"lightIM/rpc/online/online"
)

type ServiceContext struct {
	C            *config.Config
	ConnPool     *imnet.ConnPool
	OnlineRpc    online.Online
	ChatProducer *mq.ImProducer
	MessageRpc   message.Message
}

func NewServiceContext(c *config.Config) *ServiceContext {
	writerMq := mq2.NewWriterSync(&c.Edge.KqWriter)
	connOnline := zrpc.MustNewClient(c.OnlineRpc)
	connMessage := zrpc.MustNewClient(c.MessageRpc)
	return &ServiceContext{
		C:            c,
		ConnPool:     imnet.NewConnPool(context.TODO(), nil),
		OnlineRpc:    online.NewOnline(connOnline),
		ChatProducer: mq.NewImProducer(writerMq),
		MessageRpc:   message.NewMessage(connMessage),
	}
}
