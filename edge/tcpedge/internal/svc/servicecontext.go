package svc

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/common/mq"
	"lightIM/edge/tcpedge/internal/config"
	"lightIM/edge/tcpedge/internal/imnet"
	"lightIM/rpc/online/online"
)

type ServiceContext struct {
	C         *config.Config
	ConnPool  *imnet.ConnPool
	OnlineRpc online.Online
	ChatMq    *mq.Reader
}

func NewServiceContext(c *config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(c.OnlineRpc)
	return &ServiceContext{
		C:         c,
		ConnPool:  imnet.NewConnPool(context.TODO(), nil),
		OnlineRpc: online.NewOnline(conn),
		ChatMq:    mq.NewReader(&c.Edge.Kq),
	}
}
