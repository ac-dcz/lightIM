package offline

import (
	"lightIM/edge/tcpedge/internal/logic/offline"
	"lightIM/edge/tcpedge/internal/svc"
	"lightIM/edge/tcpedge/types"
)

func HandleOffline(svcCtx *svc.ServiceContext, msg *types.OfflineNotify, _ string) {
	logic := offline.NewOfflineLogic(svcCtx)
	_ = logic.Offline(msg)
	//TODO: what? just notify
}
