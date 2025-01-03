package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	BizRedis        redis.RedisConf
	UserRpc         zrpc.RpcClientConf
	MessageRpc      zrpc.RpcClientConf
	RelationshipRpc zrpc.RpcClientConf
	Auth            struct {
		AccessSecret string
		AccessExpire int64
	}
}
