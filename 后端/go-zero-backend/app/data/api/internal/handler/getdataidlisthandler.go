package handler

import (
	"net/http"

	"backend/app/data/api/internal/logic"
	"backend/app/data/api/internal/svc"
	"backend/app/data/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDataIdListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDataIdListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetDataIdListLogic(r.Context(), svcCtx)
		resp, err := l.GetDataIdList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
