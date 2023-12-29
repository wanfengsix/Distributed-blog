package search

import (
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/logic/search"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Prefix_Search_managing() http.HandlerFunc {
	return base.Prefix_Managing
}

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "GET" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			//return
		}
		reso := search.NewSearchLogic(r.Context(), svcCtx)
		resp, err := reso.Search(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
