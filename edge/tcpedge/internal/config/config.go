package config

import (
	"lightIM/common/params"
	"lightIM/common/sd"
)

type KafkaConf struct {
	Brokers   []string `json:"brokers"`
	Topic     string   `json:"topic"`
	GroupName string   `json:"groupName"`
}

type EdgeTcpServerConf struct {
	Host   string
	Key    string    //Etcd key
	EdgeId int64     //Edge id
	Kq     KafkaConf //kafka
}

func (e *EdgeTcpServerConf) EtcdKey() string {
	return e.Key
}

func (e *EdgeTcpServerConf) MetaData() sd.MetaData {
	meta := make(sd.MetaData)
	meta[params.EdgeTcpServer.EtcdEdgeId] = e.EdgeId
	meta[params.EdgeTcpServer.EtcdEdgeKq] = e.Kq
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
}
