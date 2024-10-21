package mq

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/params"
)

type pair struct {
	Msg      *kafka.Message
	CallBack func(msg *kafka.Message, err error)
}

type HandleFunc func(msg *kafka.Message) error

type MqHandlerOptions struct {
	poolSize int
}

type MqHandler struct {
	handle   HandleFunc
	workPool *ants.PoolWithFunc
}

func NewMqHandler(handle HandleFunc, opt *MqHandlerOptions) (*MqHandler, error) {
	if opt == nil {
		opt = &MqHandlerOptions{
			poolSize: params.EdgeTcpServer.WorkPoolSize,
		}
	}
	h := &MqHandler{
		handle: handle,
	}
	pool, err := ants.NewPoolWithFunc(opt.poolSize, func(i interface{}) {
		if msg, ok := i.(*pair); ok {
			var err error
			defer func() {
				if msg.CallBack != nil {
					msg.CallBack(msg.Msg, err)
				}
			}()
			err = h.handle(msg.Msg)
		}
	})
	if err != nil {
		return nil, err
	}
	h.workPool = pool
	return h, nil
}

func (h *MqHandler) HandleMessage(msg *kafka.Message, callback func(msg *kafka.Message, err error)) error {
	if err := h.workPool.Invoke(&pair{Msg: msg, CallBack: callback}); err != nil {
		logx.Errorf("mq work pool error: %v", err)
		return fmt.Errorf("handle message error: workpoll invoke fault")
	}
	return nil
}

var defaultHandler *MqHandler

func defaultHandle(msg *kafka.Message) error {
	//TODO: handle message

	return nil
}

func init() {
	var err error
	defaultHandler, err = NewMqHandler(defaultHandle, nil)
	if err != nil {
		panic(err)
	}
}
