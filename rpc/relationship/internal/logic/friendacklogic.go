package logic

import (
	"context"
	"lightIM/common/codes"
	"lightIM/common/params"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendAckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendAckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAckLogic {
	return &FriendAckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendAckLogic) FriendAck(in *types.AddFriendAck) (*types.AddFriendAckResp, error) {
	//Step: ack or not
	key := params.RpcRelationship.BizFriendReqKey(in.AddReqId)
	if from, to, e, err := l.findFriendReq(key); err != nil {
		l.Logger.Errorf("Relationship redis error: %v", err)
		return nil, err
	} else if !e {
		return &types.AddFriendAckResp{
			Base: &types.Base{
				Code: codes.RpcRelationReqExpired,
				Msg:  "请求已过期",
			},
		}, nil
	} else {
		//接受好友请求
		if in.Accept {
			//Step1: 跟新数据库
			if err := l.svcCtx.RelationShipModel.AddRelationShip(l.ctx, from, to); err != nil {
				l.Logger.Errorf("Relationship model error: %v", err)
				return nil, err
			}
		}
		//TODO:Step2: 通知from请求结果
	}

	return &types.AddFriendAckResp{}, nil
}

func (l *FriendAckLogic) findFriendReq(key string) (from, to int64, e bool, err error) {
	if reqId, err := l.svcCtx.BizRedis.GetCtx(l.ctx, key); err != nil {
		return 0, 0, false, err
	} else if reqId == "" {
		return 0, 0, false, nil
	} else {
		from, to, err = ParseFriendReqId(reqId)
		e = true
		return
	}
}
