package svc

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	mq2 "lightIM/common/mq"
	"lightIM/edge/tcpedge/internal/config"
	"lightIM/edge/tcpedge/internal/imnet"
	"lightIM/edge/tcpedge/internal/mq"
	"lightIM/rpc/online/online"
)

type ServiceContext struct {
	C            *config.Config
	ConnPool     *imnet.ConnPool
	OnlineRpc    online.Online
	ChatProducer *mq.ImProducer
}

func NewServiceContext(c *config.Config) *ServiceContext {
	writerMq := mq2.NewWriterSync(&c.Edge.KqWriter)
	conn := zrpc.MustNewClient(c.OnlineRpc)
	return &ServiceContext{
		C:            c,
		ConnPool:     imnet.NewConnPool(context.TODO(), nil),
		OnlineRpc:    online.NewOnline(conn),
		ChatProducer: mq.NewImProducer(writerMq),
	}
}
