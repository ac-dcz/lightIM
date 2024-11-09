package logic

import (
	"context"
	"lightIM/common/codes"
	"lightIM/common/params"
	"lightIM/common/utils"
	"lightIM/db/models/group"
	"lightIM/db/models/groupmember"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupLogic) CreateGroup(in *types.CreateGroupReq) (*types.CreateGroupResp, error) {
	//Step1: 创建Group
	data := &group.Group{
		GroupName: in.GroupName,
		Desc:      in.GroupDesc,
		Owner:     in.From,
	}
	data.GroupNum = utils.GenRandomNumString(params.RpcRelationship.GroupNumberLen)
	result, err := l.svcCtx.GroupModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Errorf("CreateGroup error: %v", err)
		return nil, err
	}
	data.Gid, _ = result.LastInsertId()

	//Step2: owner加入group
	info := &groupmember.GroupMember{
		Gid:    data.Gid,
		Member: data.Owner,
		Type:   params.GroupOwnerType,
	}
	if _, err := l.svcCtx.GroupMemberModel.Insert(l.ctx, info); err != nil {
		l.Logger.Errorf("CreateGroup error: %v", err)
		return nil, err
	}

	return &types.CreateGroupResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
		Gid:      data.Gid,
		GroupNum: []byte(data.GroupNum),
	}, nil
}
