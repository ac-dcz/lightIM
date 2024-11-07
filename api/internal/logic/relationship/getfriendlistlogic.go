package relationship

import (
	"context"
	"lightIM/common/codes"
	"lightIM/rpc/relationship/relationship"
	"lightIM/rpc/user/user"

	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 请求好友列表
func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// todo: add your logic here and delete this line
	if r, e := l.svcCtx.RelationshipRpc.FriendList(l.ctx, &relationship.FriendListReq{
		From: req.From,
	}); e != nil {
		l.Logger.Errorf("rpc call relationship error: %v", e)
		return &types.FriendListResp{
			Base: types.Base{
				Code: codes.InternalServerError.Code,
				Msg:  codes.InternalServerError.Msg,
			},
		}, nil
	} else if r.Base.Code != codes.OK.Code {
		l.Logger.Errorf("rpc call relationship resp: %#v", r)
		return &types.FriendListResp{
			Base: types.Base{
				Code: r.Base.Code,
				Msg:  r.Base.Msg,
			},
		}, nil
	} else {

		//Step2: get user info
		if userResp, e := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.UserInfoReq{
			UidList: r.FriendList,
		}); e != nil {
			l.Logger.Errorf("rpc call user error: %v", e)
			return &types.FriendListResp{
				Base: types.Base{
					Code: codes.InternalServerError.Code,
					Msg:  codes.InternalServerError.Msg,
				},
			}, nil
		} else if r.Base.Code != codes.OK.Code {
			l.Logger.Errorf("rpc call relationship resp: %#v", r)
			return &types.FriendListResp{
				Base: types.Base{
					Code: r.Base.Code,
					Msg:  r.Base.Msg,
				},
			}, nil
		} else {
			var friendList []types.UserInfo
			for _, info := range userResp.InfoList {
				friendList = append(friendList, types.UserInfo{
					Uid:      info.Uid,
					Tel:      info.Tel,
					NickName: info.NickName,
				})
			}
			resp = &types.FriendListResp{
				Base: types.Base{
					Code: codes.OK.Code,
					Msg:  codes.OK.Msg,
				},
				Friends: friendList,
			}
		}
	}

	return
}
