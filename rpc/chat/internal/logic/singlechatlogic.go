package logic

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"lightIM/common/codes"
	"lightIM/common/mq"
	"lightIM/common/params"
	"lightIM/rpc/chat/internal/svc"
	"lightIM/rpc/chat/types"
	"lightIM/rpc/message/mqtypes"
	"lightIM/rpc/online/online"

	"github.com/zeromicro/go-zero/core/logx"
)

type SingleChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSingleChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SingleChatLogic {
	return &SingleChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SingleChatLogic) SingleChat(in *types.SingleChatReq) (*types.SingleChatResp, error) {
	//Step1: 获取路由信息
	if resp, err := l.svcCtx.OnlineRpc.GetRoute(l.ctx, &online.RouteReq{
		UserId: in.To,
	}); err != nil {
		l.Logger.Errorf("Call online rpc fail: %v,resp: %#v", err, resp)
		return &types.SingleChatResp{
			Base: &types.Base{
				Code: codes.InternalServerError.Code,
				Msg:  err.Error(),
			},
		}, nil
	} else if resp.Base.Code != codes.OK.Code {
		l.Logger.Errorf("Call online rpc fail resp: %#v", resp)
		return &types.SingleChatResp{
			Base: &types.Base{
				Code: resp.Base.Code,
				Msg:  resp.Base.Msg,
			},
		}, nil
	} else {
		rdConf := &mq.ReaderConf{}
		if err := json.Unmarshal(resp.KqInfo, rdConf); err == nil {
			m := &mqtypes.Message{
				MsgId:        in.MsgId,
				Type:         params.ContentType(in.Type),
				Status:       params.UnRead,
				From:         in.From,
				To:           in.To,
				IsGroup:      false,
				EncodingType: in.EncodingType,
				Data:         in.Data,
				TimeStamp:    in.TimeStamp,
			}
			value, err := m.Encode()
			if err != nil {
				l.Logger.Errorf("kafka msg encode error: %v", err)
				return nil, err
			}
			//Step4: 向对应的Mq发送消息
			pLogic := NewProduceLogic(l.svcCtx)
			if err := pLogic.Produce(l.ctx, rdConf, kafka.Message{
				Key:   []byte(params.MqChatMessage),
				Value: value,
			}); err != nil {
				l.Logger.Errorf("kafka produce message error: %v", err)
				return nil, err
			}
		} else {
			return &types.SingleChatResp{
				Base: &types.Base{
					Code: codes.RpcChatParseRoute,
					Msg:  err.Error(),
				},
			}, nil
		}
	}

	return &types.SingleChatResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}
