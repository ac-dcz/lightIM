package message

import (
	"context"
	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取群历史记录
func NewGetHistoryGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryGroupLogic {
	return &GetHistoryGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHistoryGroupLogic) GetHistoryGroup(req *types.HistoryGroupReq) (resp *types.HistoryGroupResp, err error) {
	// todo: add your logic here and delete this line
	//if msgResp, e := l.svcCtx.MessageRpc.GetHistory(l.ctx, &message.HistoryReq{
	//	From: req,
	//}); e != nil {
	//	return nil, e
	//}

	return
}
