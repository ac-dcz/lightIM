package logic

import (
	"context"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberListLogic {
	return &GroupMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupMemberListLogic) GroupMemberList(in *types.GroupMemberListReq) (*types.GroupMemberListResp, error) {
	// todo: add your logic here and delete this line

	return &types.GroupMemberListResp{}, nil
}
