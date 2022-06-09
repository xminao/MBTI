package handler

import (
	"net/http"

	"backend/app/university/api/internal/logic"
	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMajorListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMajorListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetMajorListLogic(r.Context(), svcCtx)
		resp, err := l.GetMajorList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
