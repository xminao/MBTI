package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"question/api/internal/logic"
	"question/api/internal/svc"
	"question/api/internal/types"
)

func GetQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetQuestionLogic(r.Context(), svcCtx)
		resp, err := l.GetQuestion(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
