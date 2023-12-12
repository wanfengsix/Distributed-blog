package comment

import (
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/logic/comment"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Prefix_comment_managing() http.HandlerFunc {
	return base.Prefix_Managing
}

func CommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "GET" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		var req types.CommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		reso := comment.NewCommentLogic(r.Context(), svcCtx)
		resp, err := reso.Comment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
