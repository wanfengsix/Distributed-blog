package user

import (
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/logic/user"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Prefix_resource_managing() http.HandlerFunc {
	return base.Prefix_Managing
}

func UserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "GET" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			//return
		}
		user_logic := user.NewUserLogic(r.Context(), svcCtx)
		resp, err := user_logic.User(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
