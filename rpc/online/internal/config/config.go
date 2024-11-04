package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/common/mq"
)

type Config struct {
	zrpc.RpcServerConf
	BizRedis       redis.RedisConf
	EdgeEtcdHost   []string
	KqOnlineWriter mq.WriterConf
}
