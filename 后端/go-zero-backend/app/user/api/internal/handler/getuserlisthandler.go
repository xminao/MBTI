package handler

import (
	"net/http"

	"backend/app/user/api/internal/logic"
	"backend/app/user/api/internal/svc"
	"backend/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getuserlistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetuserlistLogic(r.Context(), svcCtx)
		resp, err := l.Getuserlist(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
