package resource

import (
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/logic/resource"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Prefix_resource_managing() http.HandlerFunc {
	return base.Prefix_Managing
}

func ResourceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "GET" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		var req types.ResourceReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			//return
		}
		reso := resource.NewResourceLogic(r.Context(), svcCtx)
		resp, err := reso.Resource(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
func ResourceHandler_POST(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "POST" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		var req types.ResourceReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			//return
		}
		reso := resource.NewResourceLogic(r.Context(), svcCtx)
		resp, err := reso.ResourcePOST(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
