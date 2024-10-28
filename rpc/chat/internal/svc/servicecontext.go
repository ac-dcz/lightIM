package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/rpc/chat/internal/config"
	"lightIM/rpc/chat/internal/mq"
	"lightIM/rpc/online/online"
)

type ServiceContext struct {
	Config    config.Config
	OnlineRpc online.Online
	Producer  *mq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(c.OnlineRpc)
	return &ServiceContext{
		Config:    c,
		OnlineRpc: online.NewOnline(conn),
		Producer:  mq.NewProducer(),
	}
}
