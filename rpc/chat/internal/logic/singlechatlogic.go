package logic

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"lightIM/common/codes"
	"lightIM/common/mq"
	"lightIM/common/params"
	"lightIM/common/sd"
	"lightIM/rpc/message/mqtypes"
	"lightIM/rpc/online/online"
	"strconv"

	"lightIM/rpc/chat/internal/svc"
	"lightIM/rpc/chat/types"

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
		l.Logger.Errorf("Call online rpc fail: %v", err)
		return &types.SingleChatResp{
			Base: &types.Base{
				Code: codes.RpcChatParseRoute,
				Msg:  err.Error(),
			},
		}, nil
	} else {
		//Step2: 根据路由信息到ETCD上订阅Edge-MQ信息
		if sub, err := sd.NewSubscriber(l.svcCtx.Config.EdgeEtcdHosts, resp.EdgeEtcdKey); err != nil {
			l.Logger.Errorf("NewSubscriber fail: %v", err)
			return &types.SingleChatResp{
				Base: &types.Base{
					Code: codes.RpcChatParseRoute,
					Msg:  err.Error(),
				},
			}, nil
		} else {
			//Step3: 解析Mq信息
			edgeId := strconv.FormatInt(resp.EdgeId, 10)
			datas := sub.Values()
			for _, data := range datas {
				if eId, ok := data[params.EdgeTcpServer.EtcdEdgeId].(string); ok && eId == edgeId {
					if conf, ok := data[params.EdgeTcpServer.EtcdEdgeKq]; ok {
						if byt, err := json.Marshal(conf); err == nil {
							rdConf := &mq.ReaderConf{}
							if err := json.Unmarshal(byt, rdConf); err == nil {
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
									Value: value,
								}); err != nil {
									l.Logger.Errorf("kafka produce message error: %v", err)
									return nil, err
								}
							} else {
								l.Logger.Errorf("reader kafka config error: %v", err)
								return nil, err
							}
						} else {
							l.Logger.Errorf("reader kafka config error: %v", err)
							return nil, err
						}
					} else {
						l.Logger.Errorf("kafka config not found")
						return &types.SingleChatResp{
							Base: &types.Base{
								Code: codes.RpcChatParseRoute,
								Msg:  "kafka config not found",
							},
						}, nil
					}
					break
				}
			}
		}
	}

	return &types.SingleChatResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}
