package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/api/internal/logic"
	"user/api/internal/svc"
	"user/api/internal/types"
)

func signInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSignInLogic(r.Context(), svcCtx)
		resp, err := l.SignIn(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
