package svc

import (
	"lightIM/edge/tcpedge/internal/config"
	"lightIM/edge/tcpedge/internal/handler"
	"lightIM/edge/tcpedge/internal/imnet"
)

type ServiceContext struct {
	C         *config.Config
	ConnPool  *imnet.ConnPool
	ImHandler handler.Interface
}
