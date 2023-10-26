package handler

import (
	"net/http"

	"distributedBlog/internal/logic"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DistributedBlogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDistributedBlogLogic(r.Context(), svcCtx)
		resp, err := l.DistributedBlog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
