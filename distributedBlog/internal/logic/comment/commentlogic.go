package comment

import (
	"context"
	"distributedBlog/internal/const_values"
	"distributedBlog/internal/models"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"log"
	"math/rand"
	"strconv"
	"time"

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
	//获取用户ID,根据请求用户名查询
	var A_list []*models.User_Total
	var UID string
	query := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where u_name=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &A_list, query, req.U_name)
	if err != nil {
		log.Println(err)

	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(A_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find user!"
		return
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find user!"
		UID = A_list[0].UID.String
	}
	//先更新用户评论数
	user_query := "UPDATE user SET comment_nums = comment_nums + 1 WHERE UID = ?"
	_, err1 := mysqlDB.ExecCtx(context.Background(), user_query, UID)
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
	comment_query := "INSERT INTO comment VALUES(?,?,?,?,?)"
	mysqlDateFormat := "2006-01-02"
	time_now := time.Now().Format(mysqlDateFormat)
	_, err1 = mysqlDB.ExecCtx(context.Background(), comment_query, req.Comment_ID, req.Comment_content, req.Article_ID, UID, time_now)
	if err1 != nil {
		log.Println(err1)
		return
	}
	//插入notice
	var Article_L []*models.ArticleResource
	var article models.ArticleResource
	var author models.User_Total
	var you models.User_Total
	var author_L []*models.User_Total
	var you_L []*models.User_Total
	query = "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url,abstract,is_visible from  article where Article_ID=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &Article_L, query, req.Article_ID)
	if err != nil {
		log.Println(err)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(Article_L) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find article!"
		return
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find article!"
		article = *Article_L[0]
	}
	//根据文章UID找作者
	query = "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where UID=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &author_L, query, article.UID)
	if err != nil {
		log.Println(err)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(author_L) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find user!"
		return
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find user!"
		author = *author_L[0]
	}
	query = "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where u_name=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &you_L, query, req.U_name)
	if err != nil {
		log.Println(err)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(you_L) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find user!"
		return
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find user!"
		you = *you_L[0]
	}
	//生成ID
	time_now = time.Now().Format(mysqlDateFormat)
	ID := you.UID.String + time_now + strconv.Itoa(rand.Intn(1000))
	//开始插入

	notice_query := "INSERT INTO notice VALUES(?,?,?,?,?,?,?,?,?)"
	_, err1 = mysqlDB.ExecCtx(context.Background(), notice_query, you.UID.String, you.U_name.String, author.UID.String, author.U_name.String, article.Article_ID.String, article.Head.String, "评论", 0, ID)
	if err1 != nil {
		log.Println(err1)
		return
	}
	resp.Success = true
	resp.Message = "comment success!"
	resp.Code = 200
	return
}
