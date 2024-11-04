package main

import (
	"flag"
	"fmt"
	"lightIM/rpc/message/internal/config"
	"lightIM/rpc/message/internal/mq"
	"lightIM/rpc/message/internal/server"
	"lightIM/rpc/message/internal/svc"
	"lightIM/rpc/message/types"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/message.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	if consumer, err := mq.NewConsumer(ctx, nil); err != nil {
		panic(err)
	} else {
		go consumer.Start()
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		types.RegisterMessageServer(grpcServer, server.NewMessageServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
