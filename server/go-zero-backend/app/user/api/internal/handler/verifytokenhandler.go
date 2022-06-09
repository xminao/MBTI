package handler

import (
	"net/http"

	"backend/app/user/api/internal/logic"
	"backend/app/user/api/internal/svc"
	"backend/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VerifyTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVerifyTokenLogic(r.Context(), svcCtx)
		resp, err := l.VerifyToken(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
