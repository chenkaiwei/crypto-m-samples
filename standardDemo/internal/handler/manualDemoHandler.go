package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"standardDemo/internal/logic"
	"standardDemo/internal/svc"
	"standardDemo/internal/types"
)

func ManualDemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StandardMsg
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewManualDemoLogic(r.Context(), svcCtx)
		resp, err := l.ManualDemo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
