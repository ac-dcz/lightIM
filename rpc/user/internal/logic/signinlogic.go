package logic

import (
	"context"
	"errors"
	"lightIM/common/codes"
	"lightIM/common/utils"
	usermodel "lightIM/db/models/user"

	"lightIM/rpc/user/internal/svc"
	"lightIM/rpc/user/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInLogic) SignIn(in *types.SignInReq) (*types.SignInResp, error) {
	// todo: add your logic here and delete this line
	// 检查用户是否存在
	if user, err := l.svcCtx.UserModel.FindOneByTel(l.ctx, in.Tel); errors.Is(err, usermodel.ErrNotFound) {
		return &types.SignInResp{
			Base: &types.Base{
				Code: codes.RpcUserInvalid,
				Msg:  "账号不存在或密码错误",
			},
		}, nil
	} else if err != nil {
		l.Logger.Errorf("mysql error: %v", err)
		return nil, err
	} else {
		if user.Password != utils.EncString(in.Pwd) {
			return &types.SignInResp{
				Base: &types.Base{
					Code: codes.RpcUserInvalid,
					Msg:  "账号不存在或密码错误",
				},
			}, nil
		}
		return &types.SignInResp{
			Base: &types.Base{
				Code: codes.OK.Code,
				Msg:  codes.OK.Msg,
			},
			Uid: user.Uid,
		}, nil
	}
}
