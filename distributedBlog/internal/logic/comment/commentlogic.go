package comment

import (
	"context"
	"distributedBlog/internal/const_values"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

/*
定义评论逻辑体
*/
type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var mysqlDB = sqlx.NewSqlConn("mysql", const_values.MYSQLCONNECTION)

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
评论事务:
用户在前端以指定UID，文章ID，评论内容，评论ID（目前默认UID+日期）向后端发送请求，
后端分别对用户表添加评论数，文章表添加评论数，评论表添加评论记录
*/

func (c *CommentLogic) Comment(req *types.CommentReq) (resp *types.CommentResponse, err error) {
	res := new(types.CommentResponse)
	resp = res
	//先更新用户评论数
	user_query := "UPDATE user SET comment_nums = comment_nums + 1 WHERE UID = ?"
	_, err1 := mysqlDB.ExecCtx(context.Background(), user_query, req.UID)
	if err1 != nil {
		log.Println(err)
		return
	}
	//在更新文章评论数
	article_query := "UPDATE article SET comment_nums = comment_nums + 1 WHERE Article_ID = ?"
	_, err1 = mysqlDB.ExecCtx(context.Background(), article_query, req.Article_ID)
	if err1 != nil {
		log.Println(err)
		return
	}
	//最后插入评论
	comment_query := "INSERT INTO comment VALUES(?,?,?,?)"
	_, err1 = mysqlDB.ExecCtx(context.Background(), comment_query, req.Comment_ID, req.Comment_content, req.Article_ID, req.UID)
	if err1 != nil {
		log.Println(err)
		return
	}
	resp.Success = true
	resp.Message = "comment success!"
	resp.Code = 200
	return
}
