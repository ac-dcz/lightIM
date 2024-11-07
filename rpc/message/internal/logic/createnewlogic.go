package logic

import (
	"context"
	"lightIM/common/codes"

	"lightIM/rpc/message/internal/svc"
	"lightIM/rpc/message/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateNewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNewLogic {
	return &CreateNewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateNew 当有新用户注册时，为其创建一个新的消息项
func (l *CreateNewLogic) CreateNew(in *types.CreateNewReq) (*types.CreateNewResp, error) {
	if err := l.svcCtx.HistoryModel.CreateNew(l.ctx, in.Id, in.IsGroup); err != nil {
		l.Logger.Errorf("CreateNew err:%v", err)
		return &types.CreateNewResp{
			Base: &types.Base{
				Code: codes.InternalServerError.Code,
				Msg:  codes.InternalServerError.Msg,
			},
		}, err
	}

	return &types.CreateNewResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}
