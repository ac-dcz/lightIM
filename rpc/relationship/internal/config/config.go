package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/common/mq"
)

type Config struct {
	zrpc.RpcServerConf
	RelationShipDSN string
	GroupDSN        string
	GroupMemberDSN  string
	CacheConf       cache.CacheConf
	BizRedisConf    redis.RedisConf
	OnlineRpc       zrpc.RpcClientConf
	KqOnlineReader  mq.ReaderConf
}
