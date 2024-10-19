package access

import (
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/edge/tcpedge/internal/logic/access"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
	"time"
)

func HandleAccessMsg(svcCtx *svc.ServiceContext, msg *types.AccessMsg, key string) {
	logic := access.NewAccessLogic(svcCtx)

	if conn, ok := svcCtx.ConnPool.GetUnAuthConnByAddr(key); ok {
		resp, err := logic.Access(msg)
		if err != nil {
			resp = &types.AccessMsgResp{
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

		if err = conn.Write(resp); err != nil {
			logx.Error(err)
		}
	}

}
