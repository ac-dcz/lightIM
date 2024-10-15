package login

import (
	"context"
	"lightIM/common/codes"
	"lightIM/common/jwt"
	"lightIM/common/params"
	usertypes "lightIM/rpc/user/types"

	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录
func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (resp *types.SignInResp, err error) {
	// todo: add your logic here and delete this line
	if req.Tel == "" || req.Pwd == "" {
		return &types.SignInResp{
			Base: types.Base{
				Code: codes.LoginParamsInvaild,
				Msg:  "无效参数",
			},
		}, nil
	}

	if rpcResp, err := l.svcCtx.UserRpc.SignIn(l.ctx, &usertypes.SignInReq{
		Tel: req.Tel,
		Pwd: req.Pwd,
	}); err != nil {
		l.Logger.Errorf("user rpc error: %v", err)
		return &types.SignInResp{
			Base: internalErrorBase,
		}, nil
	} else if rpcResp.Base.Code != codes.OK.Code {
		return &types.SignInResp{
			Base: types.Base{
				Code: rpcResp.Base.Code,
				Msg:  rpcResp.Base.Msg,
			},
		}, nil
	} else {
		payload := map[string]interface{}{
			params.UserIdKey: rpcResp.Uid,
		}
		if token, err := jwt.BuildToken(&jwt.TokenOption{
			AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
			AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		}, payload); err != nil {
			l.Logger.Errorf("jwt build token error: %v", err)
			return &types.SignInResp{
				Base: internalErrorBase,
			}, nil
		} else {
			resp = &types.SignInResp{
				Base: types.Base{
					Code: codes.OK.Code,
					Msg:  codes.OK.Msg,
				},
				Token: token,
			}
		}
	}
	return
}
