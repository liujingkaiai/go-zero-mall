package handler

import (
	"net/http"

	"github.com/liujingkaiai/go-zero-mall/apps/product/admin/internal/logic"
	"github.com/liujingkaiai/go-zero-mall/apps/product/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadImageLogic(r.Context(), svcCtx, r)
		resp, err := l.UploadImage()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
