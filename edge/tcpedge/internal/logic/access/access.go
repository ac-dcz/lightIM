package access

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
)

type AccessLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccessLogic(svcCtx *svc.ServiceContext) *AccessLogic {
	return &AccessLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *AccessLogic) Access(msg *types.AccessMsg) (*types.AccessMsgResp, error) {

	return &types.AccessMsgResp{}, nil
}
