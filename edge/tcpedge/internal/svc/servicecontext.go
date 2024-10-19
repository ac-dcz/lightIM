package svc

import (
	"context"
	"lightIM/edge/tcpedge/internal/config"
	"lightIM/edge/tcpedge/internal/imnet"
)

type ServiceContext struct {
	C        *config.Config
	ConnPool *imnet.ConnPool
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		C:        c,
		ConnPool: imnet.NewConnPool(context.TODO(), nil),
	}
}
