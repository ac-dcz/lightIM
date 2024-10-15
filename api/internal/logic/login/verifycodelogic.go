package login

import (
	"context"
	"fmt"
	"lightIM/common/codes"
	"lightIM/common/params"
	"lightIM/common/utils"
	"strconv"

	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var internalErrorBase = types.Base{
	Code: codes.InternalServerError.Code,
	Msg:  codes.InternalServerError.Msg,
}

type VerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewVerifyCodeLogic 请求验证码
func NewVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCodeLogic {
	return &VerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeLogic) VerifyCode(req *types.VerifyCodeReq) (*types.VerifyCodeResp, error) {
	// todo: add your logic here and delete this line
	l.Logger.Debugf("Tel: %s", req.Tel)
	if req.Tel == "" {
		return &types.VerifyCodeResp{
			Base: types.Base{
				Code: codes.LoginTelEmpty,
				Msg:  "Tel is empty",
			},
		}, nil
	}

	//限流
	if ok, sec, err := l.intervalOk(req.Tel); err != nil {
		l.Logger.Errorf("redis error %s", err)
		return &types.VerifyCodeResp{
			Base: internalErrorBase,
		}, nil
	} else if !ok {
		l.Logger.Infof("too fast %s,after %d seconds try angina", req.Tel, sec)
		return &types.VerifyCodeResp{
			Base: types.Base{
				Code: codes.LoginVerifyTooFast,
				Msg:  fmt.Sprintf("too fast %s ,after %d seconds try angina", req.Tel, sec),
			},
		}, nil
	}

	//是否超过最大次数
	if times, err := l.getReqTimes(req.Tel); err != nil {
		l.Logger.Errorf("redis error %s", err)
		return &types.VerifyCodeResp{
			Base: types.Base{
				Code: codes.InternalServerError.Code,
				Msg:  codes.InternalServerError.Msg,
			},
		}, nil
	} else if params.ApiVerifyCode.PerDayMaxTimes <= times {
		return &types.VerifyCodeResp{
			Base: types.Base{
				Code: codes.LoginVerifyOverTimes,
				Msg:  "max times",
			},
		}, nil
	} else {
		_ = l.incReqTimes(req.Tel)
	}

	//生成Code
	code := utils.GenRandomNumString(params.ApiVerifyCode.CodeLen)
	if err := l.saveCode(req.Tel, code); err != nil {
		l.Logger.Errorf("redis error %s", err)
		return &types.VerifyCodeResp{
			Base: internalErrorBase,
		}, nil
	}

	//通过短信方式将Code发送
	l.Logger.Infof("Send MS code %s", code)

	return &types.VerifyCodeResp{
		Base: types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}

func (l *VerifyCodeLogic) intervalOk(tel string) (bool, int, error) {
	intervalKey := params.ApiVerifyCode.BizIntervalKey(tel)
	if ok, err := l.svcCtx.BizRds.ExistsCtx(l.ctx, intervalKey); err != nil {
		return false, 0, err
	} else if !ok { //如果可以访问
		if ok, err := l.svcCtx.BizRds.SetnxExCtx(l.ctx, intervalKey, "", int(params.ApiVerifyCode.IntervalExpiredtime.Seconds())); err != nil {
			return false, 0, err
		} else if !ok {
			return false, 0, nil
		}
		return true, 0, nil
	} else { //间隔时间未到
		sec, _ := l.svcCtx.BizRds.TtlCtx(l.ctx, intervalKey)
		return false, sec, nil
	}
}

func (l *VerifyCodeLogic) getReqTimes(tel string) (int, error) {
	timesKey := params.ApiVerifyCode.BizTimesKey(tel)
	if ok, err := l.svcCtx.BizRds.ExistsCtx(l.ctx, timesKey); err != nil {
		return 0, err
	} else if !ok {
		if err := l.svcCtx.BizRds.SetexCtx(l.ctx, timesKey, "0", int(params.ApiVerifyCode.TimesExpiredtime.Seconds())); err != nil {
			return 0, err
		}
		return 0, nil
	} else {
		if resp, err := l.svcCtx.BizRds.GetCtx(l.ctx, timesKey); err != nil {
			return 0, err
		} else {
			times, _ := strconv.Atoi(resp)
			return times, nil
		}
	}
}

func (l *VerifyCodeLogic) incReqTimes(tel string) error {
	timesKey := params.ApiVerifyCode.BizTimesKey(tel)
	_, err := l.svcCtx.BizRds.IncrCtx(l.ctx, timesKey)
	return err
}

func (l *VerifyCodeLogic) saveCode(tel, code string) error {
	codekey := params.ApiVerifyCode.BizCodeKey(tel)
	return l.svcCtx.BizRds.SetexCtx(l.ctx, codekey, code, int(params.ApiVerifyCode.CodeExpiredTime.Seconds()))
}
