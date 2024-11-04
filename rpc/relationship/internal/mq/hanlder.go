package mq

import (
	"context"
	"github.com/panjf2000/ants/v2"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	cmq "lightIM/common/mq"
	"lightIM/common/params"
	"lightIM/rpc/relationship/internal/logic/mq"
	"lightIM/rpc/relationship/internal/svc"
	"strconv"
)

type pair struct {
	msg      *kafka.Message
	callBack cmq.CallBackFunc
}

type Handler struct {
	svcCtx   *svc.ServiceContext
	workPool *ants.PoolWithFunc
}

func NewHandler(svcCtx *svc.ServiceContext, opt *ConsumerOptions) (cmq.Handler, error) {
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

func (h *Handler) HandleMessage(msg *kafka.Message, callBack cmq.CallBackFunc) {
	r := &pair{msg: msg, callBack: callBack}
	if err := h.workPool.Invoke(r); err != nil {
		logx.Errorf("mq work pool error: %v", err)
	}
}

func (h *Handler) handle(msg *kafka.Message) error {
	switch string(msg.Key) {
	case params.MqOnlineNotify:
		uid, err := strconv.ParseInt(string(msg.Value), 10, 64)
		if err != nil {
			logx.Errorf("online mq decode msg error: %v", err)
			return err
		}
		l := mq.NewOnlineLogic(h.svcCtx)
		if err := l.Online(context.Background(), uid); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) Close() error {
	if h.workPool != nil {
		h.workPool.Release()
	}
	return nil
}
