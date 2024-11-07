package logic

import (
	"context"
	"lightIM/common/codes"

	"lightIM/rpc/user/internal/svc"
	"lightIM/rpc/user/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *types.UserInfoReq) (*types.UserInfoResp, error) {
	// todo: add your logic here and delete this line
	var infoList []*types.UserInfo
	for _, uid := range in.UidList {
		if info, err := l.svcCtx.UserModel.FindOne(l.ctx, uid); err != nil {
			l.Logger.Errorf("user model error: %v")
			return nil, err
		} else {
			infoList = append(infoList, &types.UserInfo{
				Uid:      info.Uid,
				Tel:      info.Tel,
				NickName: info.Nickname,
			})
		}
	}

	return &types.UserInfoResp{
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
		InfoList: infoList,
	}, nil
}
