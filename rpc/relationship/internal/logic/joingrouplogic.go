package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"lightIM/common/codes"
	"lightIM/common/mq"
	"lightIM/common/params"
	"lightIM/db/models/groupmember"
	"lightIM/rpc/online/online"
	"lightIM/rpc/relationship/mqtypes"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinGroupLogic) JoinGroup(in *types.JoinGroupReq) (*types.JoinGroupResp, error) {
	//Step1:  Group exists?
	group, err := l.svcCtx.GroupModel.FindOneByGroupNum(l.ctx, string(in.GroupNum))
	if errors.Is(err, groupmember.ErrNotFound) {
		return &types.JoinGroupResp{
			Base: &types.Base{
				Code: codes.RpcRelationGroupNotExists,
				Msg:  "group not exists",
			},
		}, nil
	} else if err != nil {
		l.Logger.Errorf("JoinGroup err:%v", err)
		return nil, err
	}
	//Step2: store Req in redis
	reqId := l.joinGroupReqId(group.GroupNum, in.From)
	if err := l.save(params.RpcRelationship.BizGroupReqIdKey(reqId), reqId); err != nil {
		l.Logger.Errorf("Save group req err:%v", err)
		return nil, err
	}

	//Step3: write to mq
	if resp, err := l.svcCtx.OnlineRpc.GetRoute(l.ctx, &online.RouteReq{
		UserId: group.Owner,
	}); err != nil {
		l.Logger.Errorf("Call online rpc fail: %v,resp: %#v", err, resp)
		return &types.JoinGroupResp{
			Base: &types.Base{
				Code: codes.InternalServerError.Code,
				Msg:  err.Error(),
			},
		}, nil
	} else if resp.Base.Code != codes.OK.Code {
		l.Logger.Errorf("Call online rpc fail resp: %#v", resp)
		return &types.JoinGroupResp{
			Base: &types.Base{
				Code: resp.Base.Code,
				Msg:  resp.Base.Msg,
			},
		}, nil
	} else {
		rdConf := &mq.ReaderConf{}
		if err := json.Unmarshal(resp.KqInfo, rdConf); err == nil {
			msg := &mqtypes.JoinGroupRequest{
				From:  in.From,
				Owner: group.Owner,
				Group: group.Gid,
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
					Key:   []byte(params.MqGroupReq),
					Value: value,
				}); err != nil {
					l.Logger.Errorf("Produce error: %v", err)
					return nil, err
				}

			}
		} else {
			return &types.JoinGroupResp{
				Base: &types.Base{
					Code: codes.RpcRelationParseRoute,
					Msg:  err.Error(),
				},
			}, nil
		}
	}

	return &types.JoinGroupResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}

func (l *JoinGroupLogic) joinGroupReqId(groupNumber string, from int64) string {
	return fmt.Sprintf("%s_%d", groupNumber, from)
}

func (l *JoinGroupLogic) save(key, reqId string) error {
	if err := l.svcCtx.BizRedis.SetexCtx(l.ctx, key, reqId, int(params.RpcRelationship.RdsGroupReqIdTimeout.Seconds())); err != nil {
		return err
	}
	return nil
}
