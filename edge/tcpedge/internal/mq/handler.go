package mq

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/params"
	"lightIM/edge/tcpedge/internal/logic/consumer"
	"lightIM/edge/tcpedge/internal/svc"
	msg_mq "lightIM/rpc/message/mqtypes"
	msg_req "lightIM/rpc/relationship/mqtypes"
)

type pair struct {
	Msg      *kafka.Message
	CallBack func(msg *kafka.Message, err error)
}

type consumerHandler struct {
	svcCtx   *svc.ServiceContext
	workPool *ants.PoolWithFunc
}

func newConsumerHandler(svcCtx *svc.ServiceContext, opt *ConsumerOptions) (*consumerHandler, error) {
	if opt == nil {
		opt = &ConsumerOptions{
			PoolSize: params.EdgeTcpServer.MqWorkPoolSize,
		}
	}
	h := &consumerHandler{
		svcCtx: svcCtx,
	}

	//work pool
	pool, err := ants.NewPoolWithFunc(opt.PoolSize, func(i interface{}) {
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

func (h *consumerHandler) HandleMessage(msg *kafka.Message, callback func(msg *kafka.Message, err error)) error {
	if err := h.workPool.Invoke(&pair{Msg: msg, CallBack: callback}); err != nil {
		logx.Errorf("mq work pool error: %v", err)
		return fmt.Errorf("handle message error: workpoll invoke fault")
	}
	return nil
}

func (h *consumerHandler) handle(msg *kafka.Message) error {
	switch string(msg.Key) {
	case params.MqChatMessage:
		m := &msg_mq.Message{}
		if err := m.Decode(msg.Value); err != nil {
			logx.Errorf("decode mq message error: %v", err)
			return err
		}
		l := consumer.NewLogic(h.svcCtx)
		if err := l.Exec(m); err != nil {
			logx.Errorf("logic exec error: %v", err)
			return err
		}
	case params.MqFriendReq:
		m := &msg_req.AddFriendRequest{}
		if err := m.Decode(msg.Value); err != nil {
			logx.Errorf("decode mq addfriend error: %v", err)
			return err
		}
		l := consumer.NewRequestLogic(h.svcCtx)
		if err := l.ExecFriendReq(m); err != nil {
			logx.Errorf("logic exec error: %v", err)
			return err
		}
	case params.MqGroupReq:
		m := &msg_req.JoinGroupRequest{}
		if err := m.Decode(msg.Value); err != nil {
			logx.Errorf("decode mq joingroup error: %v", err)
			return err
		}
		l := consumer.NewRequestLogic(h.svcCtx)
		if err := l.ExecGroupReq(m); err != nil {
			logx.Errorf("logic exec error: %v", err)
			return err
		}
	}

	return nil
}

func (h *consumerHandler) Close() error {
	if h.workPool != nil {
		h.workPool.Release()
	}
	return nil
}
