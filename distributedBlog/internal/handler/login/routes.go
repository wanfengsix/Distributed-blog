package login

import (
	"net/http"

	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/login",
				Handler: loginHandler(&types.LoginTokenInfo{}),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodOptions,
				Path:    "/login",
				Handler: base.Prefix_Managing,
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodOptions,
				Path:    "/user/login",
				Handler: Prefix_login_managing(),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: PostHandler(),
			},
		},
	)

}
