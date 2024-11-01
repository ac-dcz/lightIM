package logic

import (
	"context"
	"lightIM/common/codes"

	"lightIM/rpc/relationship/internal/svc"
	"lightIM/rpc/relationship/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *types.FriendListReq) (*types.FriendListResp, error) {
	// todo: add your logic here and delete this line
	if data, err := l.svcCtx.RelationShipModel.RelationshipList(l.ctx, in.From); err != nil {
		l.Logger.Errorf("relationship query error: %v", err)
		return nil, err
	} else {
		resp := &types.FriendListResp{
			Base: &types.Base{
				Code: codes.OK.Code,
				Msg:  codes.OK.Msg,
			},
		}
		for _, v := range data {
			resp.FriendList = append(resp.FriendList, v.Uid2+v.Uid1-in.From)
		}
		return resp, nil
	}
}
