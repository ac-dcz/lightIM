package handler

import (
	"github.com/panjf2000/ants/v2"
	"lightIM/edge/tcpedge/internal/logic"
)

type Request struct {
	Key string
	R   any
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

func MustNewImHandler(opt *ImHandlerOptions) *ImHandler {
	if opt == nil {
		opt = &ImHandlerOptions{
			poolSize:    10,
			reqChanSize: 100,
		}
	}
	pool, err := ants.NewPoolWithFunc(opt.poolSize, logic.HandleRequest, ants.WithNonblocking(false))
	if err != nil {
		panic(err)
	}
	imh := &ImHandler{
		workPool: pool,
		reqChan:  make(chan *Request, opt.reqChanSize),
	}

	go func() {
		for req := range imh.reqChan {
			_ = imh.workPool.Invoke(req)
		}
	}()

	return imh
}

func (imh *ImHandler) Handle(req *Request) {
	imh.reqChan <- req
}
