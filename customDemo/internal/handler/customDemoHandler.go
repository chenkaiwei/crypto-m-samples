package handler

import (
	"net/http"

	"customDemo/internal/logic"
	"customDemo/internal/svc"
	"customDemo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CustomDemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SimpleMsg
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCustomDemoLogic(r.Context(), svcCtx)
		resp, err := l.CustomDemo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
