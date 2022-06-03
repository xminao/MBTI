package handler

import (
	"net/http"

	"backend/app/university_goctl/api/internal/logic"
	"backend/app/university_goctl/api/internal/svc"
	"backend/app/university_goctl/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddYearHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddYearReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddYearLogic(r.Context(), svcCtx)
		resp, err := l.AddYear(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
