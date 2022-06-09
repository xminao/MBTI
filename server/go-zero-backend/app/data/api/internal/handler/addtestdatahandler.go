package handler

import (
	"net/http"

	"backend/app/data/api/internal/logic"
	"backend/app/data/api/internal/svc"
	"backend/app/data/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddTestDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTestDataReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddTestDataLogic(r.Context(), svcCtx)
		resp, err := l.AddTestData(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
