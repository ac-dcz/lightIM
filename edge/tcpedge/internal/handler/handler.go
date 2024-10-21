package handler

import (
	"github.com/panjf2000/ants/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/params"
	"lightIM/edge/tcpedge/internal/handler/access"
	"lightIM/edge/tcpedge/internal/handler/multichat"
	"lightIM/edge/tcpedge/internal/handler/offline"
	"lightIM/edge/tcpedge/internal/handler/singlechat"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
)

type Request struct {
	RemoteAddr string
	R          any
	SvcCtx     *svc.ServiceContext
}

type Interface interface {
	Handle(req *Request)
}

type ImHandlerOptions struct {
	poolSize    int
	reqChanSize int
}

type ImHandler struct {
	workPool *ants.PoolWithFunc
	reqChan  chan *Request
}

func NewImHandler(opt *ImHandlerOptions) (*ImHandler, error) {
	if opt == nil {
		opt = &ImHandlerOptions{
			poolSize:    params.EdgeTcpServer.WorkPoolSize,
			reqChanSize: params.EdgeTcpServer.ReqChannelBuf,
		}
	}

	imh := &ImHandler{
		reqChan: make(chan *Request, opt.reqChanSize),
	}

	pool, err := ants.NewPoolWithFunc(opt.poolSize, func(i interface{}) {
		imh.handle(i)
	}, ants.WithNonblocking(false))

	if err != nil {
		return nil, err
	}

	imh.workPool = pool

	go func() {
		for req := range imh.reqChan {
			_ = imh.workPool.Invoke(req)
		}
	}()

	return imh, nil
}

func (imh *ImHandler) Handle(req *Request) {
	imh.reqChan <- req
}

func (imh *ImHandler) handle(v any) {
	req, ok := v.(*Request)
	if !ok {
		panic("handle message is not *request")
	}
	switch msg := req.R.(type) {

	case *types.AccessMsg:
		access.HandleAccessMsg(req.SvcCtx, msg, req.RemoteAddr)
	case *types.SingleChatMsg:
		singlechat.HandleSingleChatMsg(req.SvcCtx, msg, req.RemoteAddr)
	case *types.MultiChatMsg:
		multichat.HandleMultiChatMsg(req.SvcCtx, msg, req.RemoteAddr)
	case *types.OfflineNotify:
		offline.HandleOffline(req.SvcCtx, msg, req.RemoteAddr)
	default:
		logx.Errorf("unhandle msg %v,not found msg type", msg)
	}
}
