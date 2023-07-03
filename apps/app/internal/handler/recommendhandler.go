package handler

import (
	"net/http"

	"app/internal/logic"
	"app/internal/svc"
	"app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
