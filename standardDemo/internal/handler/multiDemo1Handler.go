package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"standardDemo/internal/logic"
	"standardDemo/internal/svc"
	"standardDemo/internal/types"
)

func MultiDemo1Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SimpleMsg
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMultiDemo1Logic(r.Context(), svcCtx)
		resp, err := l.MultiDemo1(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
