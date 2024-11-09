package logic

import (
	"context"
	"lightIM/common/codes"

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
	members, err := l.svcCtx.GroupMemberModel.FindMembersByGid(l.ctx, in.GroupId)
	if err != nil {
		l.Logger.Errorf("GroupMemberListLogic FindMembersByGid err: %v", err)
		return nil, err
	}
	var uidList []int64
	for _, member := range members {
		uidList = append(uidList, member.Member)
	}
	return &types.GroupMemberListResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
		UidList: uidList,
	}, nil
}
