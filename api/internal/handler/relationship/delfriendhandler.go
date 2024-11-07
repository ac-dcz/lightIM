package relationship

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lightIM/api/internal/logic/relationship"
	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"
)

// 删除好友
func DelFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelFriendReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := relationship.NewDelFriendLogic(r.Context(), svcCtx)
		resp, err := l.DelFriend(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
