package user

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
定义用户逻辑体
*/
type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var mysqlDB = sqlx.NewSqlConn("mysql", const_values.MYSQLCONNECTION)

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
用户请求处理事务：通过请求类型 UID|文章ID，对表user进行查询将其读入，作为
响应体返回给前端页面
*/
func (u *UserLogic) User(req *types.UserReq) (resp *types.UserResponse, err error) {
	res := new(types.UserResponse)
	resp = res
	if req.Type == "Article_ID" {
		GetUser_By_Article_ID(req, resp)
	}
	return
}

// 如果req.Type为Article_ID,那么req.Value为文章ID，根据文章ID查询UID，再根据UID返回用户信息
func GetUser_By_Article_ID(req *types.UserReq, resp *types.UserResponse) {
	var R_list []*models.ArticleResource
	//根据Article_ID查UID
	Article_query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url from article where Article_ID=?"
	err1 := mysqlDB.QueryRowsCtx(context.Background(), &R_list, Article_query, req.Value)
	if err1 != nil {
		log.Println(err1)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(R_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find Article!"
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find Article!"
	}
	//根据UID查
	UID := R_list[0].UID
	var U_list []*models.User_Total
	query := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where uid=?"
	err2 := mysqlDB.QueryRowsCtx(context.Background(), &U_list, query, UID)
	if err2 != nil {
		log.Println(err2)
	}
	resp.Data = *U_list[0]
	return
}
