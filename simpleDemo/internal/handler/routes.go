// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"simpleDemo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CryptionRequest, serverCtx.CryptionResponse},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/cryptionTest",
					Handler: CryptionTestHandler(serverCtx),
				},
			}...,
		),
	)
}
