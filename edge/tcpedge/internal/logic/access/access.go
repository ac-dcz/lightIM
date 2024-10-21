package access

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/common/jwt"
	"lightIM/common/params"
	"lightIM/edge/tcpedge/internal/imnet"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"lightIM/rpc/online/online"
)

type AccessLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccessLogic(svcCtx *svc.ServiceContext) *AccessLogic {
	return &AccessLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *AccessLogic) Access(msg *types.AccessMsg, conn *imnet.ImConn) (*types.AccessMsgResp, error) {
	tokenOpt := &jwt.TokenOption{
		AccessSecret: l.svcCtx.C.Auth.AccessSecret,
		AccessExpire: l.svcCtx.C.Auth.AccessExpire,
	}
	data, err := jwt.VerifyToken(tokenOpt, msg.Token)
	if err != nil {
		return &types.AccessMsgResp{
			RespBase: types.RespBase{
				Code: codes.EdgeAuthenticatedInvalid,
				Msg:  err.Error(),
			},
		}, nil
	}
	if uid, ok := data[params.TokenUserIdKey].(float64); !ok {
		l.Logger.Errorf("user id decode error: uid is not float64", err)
		return nil, err
	} else {
		l.Logger.Infof("Conn %s auth successfully,uid %d", conn.Key(), int64(uid))
		l.svcCtx.ConnPool.AuthConn(conn.Key(), int64(uid))

		//Online Rpc
		if resp, err := l.svcCtx.OnlineRpc.UserOnline(context.Background(), &online.UserOnlineReq{
			EdgeEtcdKey: l.svcCtx.C.Edge.EtcdKey(),
			EdgeId:      l.svcCtx.C.Edge.EdgeId,
			UserId:      int64(uid),
		}); err != nil {
			logx.Errorf("online rpc call UserOnline error: %v", err)
		} else if resp.Base.Code != codes.OK.Code {
			logx.Errorf("online rpc call UserOnline resp: {code: %d,msg: %s}", resp.Base.Code, resp.Base.Msg)
		}
	}

	return &types.AccessMsgResp{
		RespBase: types.RespBase{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}
