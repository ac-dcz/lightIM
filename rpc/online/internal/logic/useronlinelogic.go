package logic

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"lightIM/common/codes"
	"lightIM/common/params"

	"lightIM/rpc/online/internal/svc"
	"lightIM/rpc/online/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOnlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOnlineLogic {
	return &UserOnlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOnlineLogic) UserOnline(in *types.UserOnlineReq) (*types.UserOnlineResp, error) {
	if in.EdgeEtcdKey == "" || in.EdgeId <= 0 || in.UserId <= 0 {
		return &types.UserOnlineResp{
			Base: &types.Base{
				Code: codes.RpcOnlineParamsInvalid,
				Msg:  "无效参数",
			},
		}, nil
	}
	if err := l.addOnlineUser(in.EdgeId, in.UserId); err != nil {
		l.Logger.Errorf("add online user error: %v", err)
		return nil, err
	}
	if err := l.updateEdgeInfo(in.EdgeId, in.EdgeEtcdKey); err != nil {
		l.Logger.Errorf("update edge info error: %v", err)
		return nil, err
	}

	//notify other rpc server
	_ = l.svcCtx.OnlineWriter.Write(l.ctx, kafka.Message{
		Key:   []byte(params.MqOnlineNotify),
		Value: []byte(fmt.Sprintf("%d", in.UserId)),
	})

	return &types.UserOnlineResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}

func (l *UserOnlineLogic) addOnlineUser(edgeId, uId int64) error {
	l.svcCtx.EdgeCache.AddOnline(edgeId, uId)
	key := params.RpcOnline.BizEdgeOnlineKey(edgeId)
	if _, err := l.svcCtx.BizRds.Sadd(key, uId); err != nil {
		return err
	}
	return nil
}

func (l *UserOnlineLogic) updateEdgeInfo(edgeId int64, etcdKey string) error {
	if t, ok := l.svcCtx.EdgeCache.GetEtcdKey(edgeId); !ok || t != etcdKey {
		l.svcCtx.EdgeCache.UpdateEtcdKey(edgeId, etcdKey)
		key := params.RpcOnline.BizEdgeOnlineKey(edgeId)
		if err := l.svcCtx.BizRds.Set(key, etcdKey); err != nil {
			return err
		}
	}
	return nil
}
