package write

import (
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/logic/write"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Prefix_managing() http.HandlerFunc {
	return base.Prefix_Managing
}

func PostWriteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "POST" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		var req types.WriteReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			//return
		}
		write_logic := write.NewWriteLogic(r.Context(), svcCtx)
		resp, err := write_logic.Write(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
func PostDraftHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "POST" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		var req types.WriteReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			//return
		}
		write_logic := write.NewWriteLogic(r.Context(), svcCtx)
		resp, err := write_logic.Draft(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
func GetDraftHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "GET" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		var req types.WriteReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			//return
		}
		write_logic := write.NewWriteLogic(r.Context(), svcCtx)
		resp, err := write_logic.GetResource_Draft(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
