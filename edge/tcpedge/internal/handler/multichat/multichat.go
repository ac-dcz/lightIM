package multichat

import (
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/edge/tcpedge/internal/imnet"
	"lightIM/edge/tcpedge/internal/logic/multichat"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"time"
)

func HandleMultiChatMsg(svcCtx *svc.ServiceContext, msg *types.GroupChatMsg, key string) {
	logic := multichat.NewMultiChatLogic(svcCtx)
	var (
		resp *types.GroupChatMsgResp
		err  error
		conn *imnet.ImConn
		ok   bool
	)
	conn, ok = svcCtx.ConnPool.GetAuthConnByAddr(key)
	if ok {
		if resp, err = logic.MultiChat(msg, conn); err != nil {
			resp = &types.GroupChatMsgResp{
				RespBase: types.RespBase{
					Code: codes.InternalServerError.Code,
					Msg:  err.Error(),
				},
			}
		}
	} else if conn, ok = svcCtx.ConnPool.GetUnAuthConnByAddr(key); ok {
		resp = &types.GroupChatMsgResp{
			RespBase: types.RespBase{
				Code: codes.EdgeUnAuthenticated,
				Msg:  "UnAuthenticated",
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
