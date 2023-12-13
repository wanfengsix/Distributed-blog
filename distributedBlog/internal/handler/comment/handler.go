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

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "POST" {
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

// func CommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//处理跨域请求
// 		base.Prefix_Managing(w, r)
// 		if r.Method != "POST" {
// 			log.Fatalln("请求类型错误！请检查路由配置")
// 			return
// 		}

// 		var req types.CommentReq
// 		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 			// 如果解析失败，返回错误响应
// 			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
// 			return
// 		}

// 		// var req types.CommentReq
// 		// if err := httpx.Parse(r, &req); err != nil {
// 		// 	httpx.ErrorCtx(r.Context(), w, err)
// 		// 	return
// 		// }

// 		//处理评论的业务逻辑，使用 req 中的数据进行操作
// 		reso := comment.NewCommentLogic(r.Context(), svcCtx)
// 		resp, err := reso.Comment(&req)

// 		if err != nil {
// 			httpx.ErrorCtx(r.Context(), w, err)
// 		} else {
// 			httpx.OkJsonCtx(r.Context(), w, resp)
// 		}

// 	}
// }
