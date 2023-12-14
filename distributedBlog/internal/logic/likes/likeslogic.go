package likes

import (
	"context"
	"distributedBlog/internal/const_values"
	"distributedBlog/internal/models"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

/*
定义评论逻辑体
*/
type LikesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var mysqlDB = sqlx.NewSqlConn("mysql", const_values.MYSQLCONNECTION)

func NewLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikesLogic {
	return &LikesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
点赞事务:
用户在前端以指定UID，文章ID向后端发送请求，
后端分别对用户表添加点赞数，文章表添加点赞数，点赞表添加点赞记录
*/

func (l *LikesLogic) Likes(req *types.LikesReq) (resp *types.LikesResponse, err error) {
	res := new(types.LikesResponse)
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
	user_query := "UPDATE user SET likes_nums = likes_nums + 1 WHERE UID = ?"
	_, err1 := mysqlDB.ExecCtx(context.Background(), user_query, UID)
	if err1 != nil {
		log.Println(err)
		return
	}
	//在更新文章评论数
	article_query := "UPDATE article SET likes_nums = likes_nums + 1 WHERE Article_ID = ?"
	_, err1 = mysqlDB.ExecCtx(context.Background(), article_query, req.Article_ID)
	if err1 != nil {
		log.Println(err)
		return
	}
	//最后插入点赞记录
	likes_query := "INSERT INTO likes VALUES(?,?)"
	_, err1 = mysqlDB.ExecCtx(context.Background(), likes_query, UID, req.Article_ID)
	if err1 != nil {
		log.Println(err)
		return
	}
	resp.Success = true
	resp.Message = "likes success!"
	resp.Code = 200
	return
}
