package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/common/mq"
)

type Config struct {
	zrpc.RpcServerConf
	HistoryConf struct {
		MongoCache cache.CacheConf
		Uri        string
		DB         string
		Collection string
	}
	MessageConf struct {
		MongoCache cache.CacheConf
		Uri        string
		DB         string
		Collection string
	}
	MsgReader      mq.ReaderConf
	ChatRpc        zrpc.RpcClientConf
	KqOnlineReader mq.ReaderConf
}
