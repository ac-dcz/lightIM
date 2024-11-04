package mq

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/rpc/relationship/internal/svc"
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

	return nil
}
