package offline

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"lightIM/rpc/online/online"
)

type OfflineLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOfflineLogic(svcCtx *svc.ServiceContext) *OfflineLogic {
	return &OfflineLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *OfflineLogic) Offline(msg *types.OfflineNotify) error {
	if resp, err := l.svcCtx.OnlineRpc.UserOffline(context.Background(), &online.UserOfflineReq{
		EdgeId: l.svcCtx.C.Edge.EdgeId,
		UserId: msg.Uid,
	}); err != nil {
		logx.Errorf("online rpc call UserOffline error: %v", err)
		return err
	} else if resp.Base.Code != codes.OK.Code {
		logx.Errorf("online rpc call UserOffline resp: {code: %d,msg: %s}", resp.Base.Code, resp.Base.Msg)
		return errors.New(resp.Base.Msg)
	}
	return nil
}
