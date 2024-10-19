package singlechat

import (
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/edge/tcpedge/internal/imnet"
	"lightIM/edge/tcpedge/internal/logic/singlechat"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"time"
)

func HandleSingleChatMsg(svcCtx *svc.ServiceContext, msg *types.SingleChatMsg, key string) {
	logic := singlechat.NewSingleChatLogic(svcCtx)

	var (
		resp *types.SingleChatMsgResp
		err  error
		conn *imnet.ImConn
		ok   bool
	)

	conn, ok = svcCtx.ConnPool.GetAuthConnByAddr(key)
	if ok {
		if resp, err = logic.SingleChat(msg); err != nil {
			resp = &types.SingleChatMsgResp{
				RespBase: types.RespBase{
					Base: types.Base{
						MsgId:     msg.MsgId,
						TimeStamp: time.Now().Unix(),
					},
					Code: codes.InternalServerError.Code,
					Msg:  err.Error(),
				},
			}
		}
	} else if conn, ok = svcCtx.ConnPool.GetUnAuthConnByAddr(key); ok {
		resp = &types.SingleChatMsgResp{
			RespBase: types.RespBase{
				Code: codes.EdgeUnAuthenticated,
				Msg:  "UnAuthenticated",
				Base: types.Base{
					MsgId:     msg.MsgId,
					TimeStamp: time.Now().Unix(),
				},
			},
		}
	}

	if err = conn.Write(resp); err != nil {
		logx.Error(err)
	}

}
