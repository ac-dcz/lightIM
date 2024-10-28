package mq

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/params"
	"lightIM/rpc/message/internal/svc"
)

type pair struct {
	msg      *kafka.Message
	callBack func(msg *kafka.Message, err error)
}

type Handler struct {
	svcCtx   *svc.ServiceContext
	workPool *ants.PoolWithFunc
}

func NewHandler(svcCtx *svc.ServiceContext, opt *ConsumerOptions) (*Handler, error) {
	if opt == nil {
		opt = &ConsumerOptions{
			PoolSize: params.RpcMessage.MqWorkPoolSize,
		}
	}
	h := &Handler{svcCtx: svcCtx}
	pool, err := ants.NewPoolWithFunc(opt.PoolSize, func(i interface{}) {
		if r, ok := i.(*pair); ok {
			var err error
			defer func() {
				if r.callBack != nil {
					r.callBack(r.msg, err)
				}
			}()
			err = h.handle(r.msg)
		}
	})
	if err != nil {
		return nil, err
	}
	h.workPool = pool
	return h, nil
}

func (h *Handler) HandleMessage(msg *kafka.Message, callBack func(msg *kafka.Message, err error)) error {
	r := &pair{msg: msg, callBack: callBack}
	if err := h.workPool.Invoke(r); err != nil {
		logx.Errorf("mq work pool error: %v", err)
		return fmt.Errorf("handle message error: workpoll invoke fault")
	}
	return nil
}

func (h *Handler) handle(msg *kafka.Message) error {

	return nil
}

func (h *Handler) Close() error {
	if h.workPool != nil {
		h.workPool.Release()
	}
	return nil
}
