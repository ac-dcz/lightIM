package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"lightIM/common/codes"
	"lightIM/common/mq"
	"lightIM/common/params"
	"lightIM/rpc/online/online"
	"lightIM/rpc/relationship/mqtypes"
	"strings"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func GenFriendReqId(from, to int64) string {
	return fmt.Sprintf("%d_%d", to, from)
}

func ReaderConfkey(readConf *mq.ReaderConf) string {
	return fmt.Sprintf("%s#%s", strings.Join(readConf.Brokers, "-"), readConf.Topic)
}

func ParseFriendReqId(redId string) (int64, int64, error) {
	var from, to int64 = 0, 0
	_, err := fmt.Sscanf(redId, "%d_%d", &to, &from)
	return from, to, err
}

func (l *AddFriendLogic) AddFriend(in *types.AddFriendReq) (*types.AddFriendResp, error) {
	//Step1: generator reqID and Store in redis
	reqId := GenFriendReqId(in.From, in.To)
	key := params.RpcRelationship.BizFriendReqKey(reqId)
	if l.exists(key) {
		return &types.AddFriendResp{
			Base: &types.Base{
				Code: codes.RpcRelationReqExists,
				Msg:  "请求已经存在",
			},
		}, nil
	} else {
		if err := l.save(key, reqId); err != nil {
			l.Logger.Errorf("redis error: %v", err)
			return &types.AddFriendResp{
				Base: &types.Base{
					Code: codes.InternalServerError.Code,
					Msg:  codes.InternalServerError.Msg,
				},
			}, nil
		}
	}
	//Step2: get route and write to mq

	if resp, err := l.svcCtx.OnlineRpc.GetRoute(l.ctx, &online.RouteReq{
		UserId: in.To,
	}); err != nil {
		l.Logger.Errorf("Call online rpc fail: %v,resp: %#v", err, resp)
		return &types.AddFriendResp{
			Base: &types.Base{
				Code: codes.InternalServerError.Code,
				Msg:  err.Error(),
			},
		}, nil
	} else if resp.Base.Code != codes.OK.Code {
		l.Logger.Errorf("Call online rpc fail resp: %#v", resp)
		return &types.AddFriendResp{
			Base: &types.Base{
				Code: resp.Base.Code,
				Msg:  resp.Base.Msg,
			},
		}, nil
	} else {
		rdConf := &mq.ReaderConf{}
		if err := json.Unmarshal(resp.KqInfo, rdConf); err == nil {
			msg := &mqtypes.AddFriendRequest{
				From:  in.From,
				To:    in.To,
				RedId: reqId,
			}
			if value, err := msg.Encode(); err != nil {
				l.Logger.Errorf("mq message encoder error: %v", err)
				return nil, err
			} else {

				//Write to mq
				if err := l.svcCtx.Producer.Produce(l.ctx, ReaderConfkey(rdConf), mq.WriterConf{
					Brokers: rdConf.Brokers,
					Topic:   rdConf.Topic,
				}, kafka.Message{
					Key:   []byte(params.MqFriendReq),
					Value: value,
				}); err != nil {
					l.Logger.Errorf("Produce error: %v", err)
					return nil, err
				}

			}
		} else {
			return &types.AddFriendResp{
				Base: &types.Base{
					Code: codes.RpcRelationParseRoute,
					Msg:  err.Error(),
				},
			}, nil
		}
	}

	return &types.AddFriendResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}

func (l *AddFriendLogic) exists(key string) bool {
	ok, _ := l.svcCtx.BizRedis.ExistsCtx(l.ctx, key)
	return ok
}

func (l *AddFriendLogic) save(key string, reqId string) error {
	if err := l.svcCtx.BizRedis.SetexCtx(l.ctx, key, reqId, int(params.RpcRelationship.RdsFriendReqIdTimeout.Seconds())); err != nil {
		return err
	}
	return nil
}
