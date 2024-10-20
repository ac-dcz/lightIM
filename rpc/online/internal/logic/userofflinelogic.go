package logic

import (
	"context"
	"lightIM/common/codes"
	"lightIM/common/params"

	"lightIM/rpc/online/internal/svc"
	"lightIM/rpc/online/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOfflineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOfflineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOfflineLogic {
	return &UserOfflineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOfflineLogic) UserOffline(in *types.UserOfflineReq) (*types.UserOfflineResp, error) {
	// todo: add your logic here and delete this line
	if in.UserId <= 0 || in.EdgeId <= 0 {
		return &types.UserOfflineResp{
			Base: &types.Base{
				Code: codes.RpcOnlineParamsInvalid,
				Msg:  "无效参数",
			},
		}, nil
	}

	if err := l.remOnlineUser(in.EdgeId, in.UserId); err != nil {
		l.Logger.Errorf("remove online user error: %v", err)
		return nil, err
	}

	return &types.UserOfflineResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}

func (l *UserOfflineLogic) remOnlineUser(edgeId, uId int64) error {
	l.svcCtx.EdgeCache.RemOnline(edgeId, uId)
	key := params.RpcOnline.BizEdgeOnlineKey(edgeId)
	if _, err := l.svcCtx.BizRds.Srem(key, uId); err != nil {
		return err
	}
	return nil
}
