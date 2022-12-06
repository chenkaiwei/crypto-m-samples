package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"simpleDemo/internal/logic"
	"simpleDemo/internal/svc"
	"simpleDemo/internal/types"
)

func CryptionTestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SimpleMsg
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCryptionTestLogic(r.Context(), svcCtx)
		resp, err := l.CryptionTest(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
