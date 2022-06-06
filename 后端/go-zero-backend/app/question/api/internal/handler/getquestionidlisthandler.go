package handler

import (
	"net/http"

	"backend/app/question/api/internal/logic"
	"backend/app/question/api/internal/svc"
	"backend/app/question/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetQuestionIdListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetQuestionIdListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetQuestionIdListLogic(r.Context(), svcCtx)
		resp, err := l.GetQuestionIdList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
