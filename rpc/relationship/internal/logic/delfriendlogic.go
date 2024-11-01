package logic

import (
	"context"
	"errors"
	"lightIM/common/codes"
	"lightIM/db/models/relationship"

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
	//Step: delete friend from database
	if err := l.svcCtx.RelationShipModel.DelRelationShip(l.ctx, in.From, in.To); err != nil {
		if errors.Is(err, relationship.ErrNotFound) {
			return &types.DelFriendResp{
				Base: &types.Base{
					Code: codes.RpcRelationNotExists,
					Msg:  "relationship not exists",
				},
			}, nil
		}
		logx.Errorf("DelFriendLogic error: %v", err)
		return nil, err
	}

	return &types.DelFriendResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}
