package handler

import (
	"net/http"

	"backend/app/question_goctl/api/internal/logic"
	"backend/app/question_goctl/api/internal/svc"
	"backend/app/question_goctl/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteQuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteQuestionLogic(r.Context(), svcCtx)
		resp, err := l.DeleteQuestion(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
