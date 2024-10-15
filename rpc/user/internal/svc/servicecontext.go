package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	usermodel "lightIM/db/models/user"
	"lightIM/rpc/user/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel usermodel.UserInfosModel
	BizRedis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DSN)
	bizRds := redis.MustNewRedis(c.BizRedis)
	return &ServiceContext{
		Config:    c,
		UserModel: usermodel.NewUserInfosModel(conn, c.UserCache),
		BizRedis:  bizRds,
	}
}
