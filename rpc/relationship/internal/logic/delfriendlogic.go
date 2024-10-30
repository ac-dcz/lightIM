package logic

import (
	"context"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFriendLogic {
	return &DelFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFriendLogic) DelFriend(in *types.DelFriendReq) (*types.DelFriendResp, error) {
	// todo: add your logic here and delete this line

	return &types.DelFriendResp{}, nil
}
