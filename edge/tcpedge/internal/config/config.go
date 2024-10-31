package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/common/mq"
	"lightIM/common/params"
	"lightIM/common/sd"
	"strconv"
)

type EdgeTcpServerConf struct {
	Host string
	Etcd struct {
		Host []string
		Key  string //Etcd key
	}
	EdgeId          int64         //Edge id
	KqReader        mq.ReaderConf //kafka
	KqMessageWriter mq.WriterConf
}

func (e *EdgeTcpServerConf) EtcdKey() string {
	return e.Etcd.Key
}

func (e *EdgeTcpServerConf) MetaData() sd.MetaData {
	meta := make(sd.MetaData)
	meta[params.EdgeTcpServer.EtcdEdgeId] = strconv.FormatInt(e.EdgeId, 10)
	meta[params.EdgeTcpServer.EtcdEdgeKq] = e.KqReader
	meta[params.EdgeTcpServer.EtcdEdgeHost] = e.Host
	return meta
}

type Config struct {
	Edge EdgeTcpServerConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	OnlineRpc  zrpc.RpcClientConf
	MessageRpc zrpc.RpcClientConf
}
