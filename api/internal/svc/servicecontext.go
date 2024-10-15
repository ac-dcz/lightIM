package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/api/internal/config"
	userclient "lightIM/rpc/user/user"
)

type ServiceContext struct {
	Config  config.Config
	BizRds  *redis.Redis
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	bizRds := redis.MustNewRedis(c.BizRedis)
	conn := zrpc.MustNewClient(c.UserRpc)
	userRpc := userclient.NewUser(conn)
	return &ServiceContext{
		Config:  c,
		BizRds:  bizRds,
		UserRpc: userRpc,
	}
}
