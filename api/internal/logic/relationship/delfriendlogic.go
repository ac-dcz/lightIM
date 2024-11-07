package relationship

import (
	"context"
	"lightIM/common/codes"
	"lightIM/rpc/relationship/relationship"

	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除好友
func NewDelFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFriendLogic {
	return &DelFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelFriendLogic) DelFriend(req *types.DelFriendReq) (resp *types.DelFriendResp, err error) {
	// todo: add your logic here and delete this line
	if r, e := l.svcCtx.RelationshipRpc.DelFriend(l.ctx, &relationship.DelFriendReq{
		From: req.From,
		To:   req.To,
	}); e != nil {
		l.Logger.Errorf("rpc call relationship error: %v", e)
		return &types.DelFriendResp{
			Base: types.Base{
				Code: codes.InternalServerError.Code,
				Msg:  codes.InternalServerError.Msg,
			},
		}, nil
	} else if r.Base.Code != codes.OK.Code {
		l.Logger.Errorf("rpc call relationship resp: %#v", r)
		return &types.DelFriendResp{
			Base: types.Base{
				Code: r.Base.Code,
				Msg:  r.Base.Msg,
			},
		}, nil
	}

	resp = &types.DelFriendResp{
		Base: types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}

	return
}
