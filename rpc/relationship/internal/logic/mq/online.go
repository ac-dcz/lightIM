package mq

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/mq"
	"lightIM/common/params"
	"lightIM/rpc/online/online"
	"lightIM/rpc/relationship/internal/logic"
	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/mqtypes"
	"strconv"
)

type OnlineLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOnlineLogic(svcCtx *svc.ServiceContext) *OnlineLogic {
	return &OnlineLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *OnlineLogic) Online(ctx context.Context, uid int64) error {
	//TODO:将与该用户未处理的请求发送
	prefix := params.RpcRelationship.BizFriendReqKey(strconv.Itoa(int(uid)) + "*")

	if keys, err := l.svcCtx.BizRedis.Keys(prefix); err != nil {
		l.Logger.Errorf("online redis error %v", err)
		return err
	} else {
		for _, key := range keys {
			if reqId, err := l.svcCtx.BizRedis.GetCtx(ctx, key); err == nil {
				from, uid, _ := logic.ParseFriendReqId(reqId)
				if resp, err := l.svcCtx.OnlineRpc.GetRoute(context.Background(), &online.RouteReq{
					UserId: uid,
				}); err == nil {
					rdConf := &mq.ReaderConf{}
					if err := json.Unmarshal(resp.KqInfo, rdConf); err == nil {
						msg := &mqtypes.AddFriendRequest{
							From:  from,
							To:    uid,
							RedId: reqId,
						}
						if value, err := msg.Encode(); err != nil {
							l.Logger.Errorf("mq message encoder error: %v", err)
							return err
						} else {

							//Write to mq
							if err := l.svcCtx.Producer.Produce(context.Background(), logic.ReaderConfkey(rdConf), mq.WriterConf{
								Brokers: rdConf.Brokers,
								Topic:   rdConf.Topic,
							}, kafka.Message{
								Key:   []byte(params.MqFriendReq),
								Value: value,
							}); err != nil {
								l.Logger.Errorf("Produce error: %v", err)
								return err
							}

						}
					}
				}
			} else {
				l.Logger.Errorf("online redis error %v", err)
				return err
			}
		}
	}

	return nil
}
