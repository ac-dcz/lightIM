package ackmsg

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/edge/tcpedge/internal/imnet"
	"lightIM/edge/tcpedge/internal/logic/ackmsg"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"time"
)

func HandleAckMsg(svcCtx *svc.ServiceContext, msg *types.AckMsg, key string) {
	logic := ackmsg.NewAckMsgLogic(svcCtx)
	var (
		resp *types.AckMsgResp
		err  error
		conn *imnet.ImConn
		ok   bool
	)
	if conn, ok = svcCtx.ConnPool.GetAuthConnByAddr(key); ok {
		resp, err = logic.AckMsg(context.Background(), msg)
		if err != nil {
			resp = &types.AckMsgResp{
				RespBase: types.RespBase{
					Code: codes.InternalServerError.Code,
					Msg:  err.Error(),
				},
			}
		}

		resp.Base = types.Base{
			MsgId:     msg.MsgId,
			TimeStamp: time.Now().Unix(),
		}

		if err = conn.Write(resp); err != nil {
			logx.Error(err)
		}
	}

}
