package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"lightIM/common/mq"
	"reflect"
	"testing"
)

var testConfig = Config{
	Edge: EdgeTcpServerConf{
		Etcd: struct {
			Host []string
			Key  string
		}{Host: []string{"127.0.0.1:2379"}, Key: "edge.tcp"},
		EdgeId: 1001,
		Host:   "127.0.0.1:6000",
		KqReader: mq.ReaderConf{
			Topic:     "edge.tcp.1001",
			GroupName: "edge.tcp.1001.mq",
			Brokers:   []string{"127.0.0.1:9092"},
		},
	},
	Auth: struct {
		AccessSecret string
		AccessExpire int64
	}{AccessSecret: "AFEFBCDDEFC", AccessExpire: 3600},
}

func TestBuildConfig(t *testing.T) {
	cfg := &Config{}
	if err := conf.Load("test.yaml", cfg); err != nil {
		t.Error(err)
	}
	t.Logf("%#v", cfg)
	if !reflect.DeepEqual(testConfig, *cfg) {
		t.Errorf("no equal")
	}
}
