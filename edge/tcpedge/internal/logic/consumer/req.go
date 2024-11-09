package consumer

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/rpc/relationship/mqtypes"
)

type RequestLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRequestLogic(svcCtx *svc.ServiceContext) *RequestLogic {
	return &RequestLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *RequestLogic) ExecFriendReq(req *mqtypes.AddFriendRequest) error {
	return nil
}

func (l *RequestLogic) ExecGroupReq(req *mqtypes.JoinGroupRequest) error {
	return nil
}
