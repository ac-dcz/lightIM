package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/common/mq"
	"lightIM/common/params"
	"lightIM/common/sd"
)

type EdgeTcpServerConf struct {
	Host     string
	Key      string        //Etcd key
	EdgeId   int64         //Edge id
	KqReader mq.ReaderConf //kafka
	KqWriter mq.WriterConf
}

func (e *EdgeTcpServerConf) EtcdKey() string {
	return e.Key
}

func (e *EdgeTcpServerConf) MetaData() sd.MetaData {
	meta := make(sd.MetaData)
	meta[params.EdgeTcpServer.EtcdEdgeId] = e.EdgeId
	meta[params.EdgeTcpServer.EtcdEdgeKq] = e.KqReader
	meta[params.EdgeTcpServer.EtcdEdgeHost] = e.Host
	return meta
}

type Config struct {
	Edge EdgeTcpServerConf
	Etcd struct {
		Host []string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	OnlineRpc zrpc.RpcClientConf
}
