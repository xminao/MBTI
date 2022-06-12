package handler

import (
	"net/http"

	"backend/app/data/api/internal/logic"
	"backend/app/data/api/internal/svc"
	"backend/app/data/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTestDataReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetDataLogic(r.Context(), svcCtx)
		resp, err := l.GetData(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
