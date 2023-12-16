package user

import (
	"net/http"

	"distributedBlog/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodOptions,
				Path:    "/userinfo/:type/:value",
				Handler: Prefix_resource_managing(),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/userinfo/:type/:value",
				Handler: UserHandler(serverCtx),
			},
		},
	)
}
