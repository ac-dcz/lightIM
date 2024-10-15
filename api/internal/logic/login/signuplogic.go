package login

import (
	"context"
	"lightIM/common/codes"
	"lightIM/common/params"

	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	usertypes "lightIM/rpc/user/types"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewSignUpLogic 账号注册
func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.SignUpReq) (resp *types.SignUpResp, err error) {
	// todo: add your logic here and delete this line
	if req.Tel == "" || req.Pwd == "" || req.NickName == "" {
		return &types.SignUpResp{
			Base: types.Base{
				Code: codes.LoginParamsInvaild,
				Msg:  "无效参数",
			},
		}, nil
	}

	//check code
	if ok, err := l.checkCode(req.Tel, req.Code); err != nil {
		l.Logger.Errorf("redis error %v", err)
		return &types.SignUpResp{
			Base: internalErrorBase,
		}, nil
	} else if !ok {
		return &types.SignUpResp{
			Base: types.Base{Code: codes.LoginVerifyCodeErr, Msg: "验证码无效"},
		}, nil
	}

	//SignUp
	if resp, err := l.svcCtx.UserRpc.SignUp(l.ctx, &usertypes.SignUpReq{
		Tel:      req.Tel,
		Pwd:      req.Pwd,
		Nickname: req.NickName,
	}); err != nil {
		l.Logger.Errorf("user rpc error: %v", err)
		return &types.SignUpResp{
			Base: internalErrorBase,
		}, nil
	} else if resp.Base.Code != codes.OK.Code {
		return &types.SignUpResp{
			Base: types.Base{
				Code: resp.Base.Code,
				Msg:  resp.Base.Msg,
			},
		}, nil
	}

	return &types.SignUpResp{
		Base: types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}

func (l *SignUpLogic) checkCode(tel, code string) (bool, error) {
	codekey := params.ApiVerifyCode.BizCodeKey(tel)
	if data, err := l.svcCtx.BizRds.GetCtx(l.ctx, codekey); err != nil {
		return false, err
	} else if data != code {
		return false, nil
	}
	return true, nil
}
