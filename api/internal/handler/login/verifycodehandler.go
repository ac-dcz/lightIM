package login

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lightIM/api/internal/logic/login"
	"lightIM/api/internal/svc"
	"lightIM/api/internal/types"
)

// 请求验证码
func VerifyCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewVerifyCodeLogic(r.Context(), svcCtx)
		resp, err := l.VerifyCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
