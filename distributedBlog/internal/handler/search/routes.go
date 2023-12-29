package search

import (
	"net/http"

	"distributedBlog/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/search/:name",
				Handler: SearchHandler(serverCtx),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodOptions,
				Path:    "/search/:name",
				Handler: Prefix_Search_managing(),
			},
		},
	)
}
