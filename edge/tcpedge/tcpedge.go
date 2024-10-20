package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"lightIM/edge/tcpedge/internal/config"
	"lightIM/edge/tcpedge/internal/handler"
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
	imHandler := handler.MustNewImHandler(nil)

	s := server.NewTcpServer(svcCtx, imHandler)
	log.Println("start listen ", cfg.Edge.Host)
	panic(s.Start())
}
