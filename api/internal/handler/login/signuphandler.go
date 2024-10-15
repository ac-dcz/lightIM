package login

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lightIM/api/internal/logic/login"
	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"
)

// 账号注册
func SignUpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignUpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewSignUpLogic(r.Context(), svcCtx)
		resp, err := l.SignUp(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
