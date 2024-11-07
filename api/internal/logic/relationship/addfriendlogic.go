package relationship

import (
	"context"
	"lightIM/common/codes"
	"lightIM/rpc/relationship/relationship"

	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友请求
func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.AddFriendReq) (resp *types.AddFriendResp, err error) {
	// todo: add your logic here and delete this line

	if r, rerr := l.svcCtx.RelationshipRpc.AddFriend(l.ctx, &relationship.AddFriendReq{
		From: req.From,
		To:   req.To,
	}); rerr != nil {
		l.Logger.Errorf("rpc call relationship error: %v", rerr)
		return &types.AddFriendResp{
			Base: types.Base{
				Code: codes.InternalServerError.Code,
				Msg:  codes.InternalServerError.Msg,
			},
		}, nil
	} else if r.Base.Code != codes.OK.Code {
		l.Logger.Errorf("rpc call relationship resp: %#v", r)
		return &types.AddFriendResp{
			Base: types.Base{
				Code: r.Base.Code,
				Msg:  r.Base.Msg,
			},
		}, nil
	}

	resp = &types.AddFriendResp{
		Base: types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}
	return
}
