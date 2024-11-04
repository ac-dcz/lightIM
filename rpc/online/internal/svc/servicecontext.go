package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"lightIM/common/mq"
	"lightIM/rpc/online/internal/cache"
	"lightIM/rpc/online/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	BizRds       *redis.Redis
	EdgeCache    *cache.EdgeInfo
	OnlineWriter *mq.Writer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		BizRds:       redis.MustNewRedis(c.BizRedis),
		EdgeCache:    cache.NewEdgeInfo(),
		OnlineWriter: mq.NewWriterAsync(&c.KqOnlineWriter, nil),
	}
}
