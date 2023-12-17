package follow

import (
	"distributedBlog/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/follow",
				Handler: FollowHandler(serverCtx),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodOptions,
				Path:    "/follow",
				Handler: Prefix_Follow_managing(),
			},
		},
	)
}
