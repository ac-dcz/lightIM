package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	cmq "lightIM/common/mq"
	"lightIM/edge/tcpedge/internal/config"
	"lightIM/edge/tcpedge/internal/handler"
	"lightIM/edge/tcpedge/internal/mq"
	"lightIM/edge/tcpedge/internal/server"
	"lightIM/edge/tcpedge/internal/svc"
	"log"
)

var filepath = flag.String("file", "./etc/tcpedge.yaml", "filepath")

func main() {
	flag.Parse()
	cfg := &config.Config{}
	if err := conf.Load(*filepath, cfg); err != nil {
		panic(err)
	}

	svcCtx := svc.NewServiceContext(cfg)
	msgReader := cmq.NewReader(&cfg.Edge.KqReader)
	if consumer, err := mq.NewImConsumer(svcCtx, msgReader, nil); err != nil {
		panic(err)
	} else {
		go consumer.Start()
	}

	imHandler, err := handler.NewImHandler(nil)
	if err != nil {
		panic(err)
	}
	s := server.NewTcpServer(svcCtx, imHandler)
	log.Println("start listen ", cfg.Edge.Host)
	panic(s.Start())
}
