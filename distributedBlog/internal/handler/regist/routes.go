package regist

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
				Path:    "/user/regist",
				Handler: Prefix_regist_managing(),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/regist",
				Handler: PostHandler(),
			},
		},
	)

}
