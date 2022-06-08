package handler

import (
	"net/http"

	"backend/app/question/api/internal/logic"
	"backend/app/question/api/internal/svc"
	"backend/app/question/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditQuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEditQuestionLogic(r.Context(), svcCtx)
		resp, err := l.EditQuestion(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
