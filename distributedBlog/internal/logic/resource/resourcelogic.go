package resource

import (
	"context"
	"distributedBlog/internal/const_values"
	"distributedBlog/internal/models"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"encoding/base64"
	"io"
	"log"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

/*
定义资源逻辑体
*/
type ResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var mysqlDB = sqlx.NewSqlConn("mysql", const_values.MYSQLCONNECTION)

func NewResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceLogic {
	return &ResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
资源请求处理事务：通过请求类型 头像|文章，对表user_information 进行查询，查询有关的资源文件路径，将其读入，作为
响应体返回给前端页面
*/

// 先根据用户名查用户id,然后再查询
func GetResource_Avatar(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	resource_type := "avatar"
	var R_list []*models.AvatarResource
	var user []*models.User_Total
	prequery := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where u_name=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &user, prequery, req.Name)
	if err != nil {
		log.Println(err)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(user) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find user!"
		return
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find user!"
	}
	query := "select uid,signature,avatar_url from user_information where uid=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, user[0].UID)
	if err != nil {
		log.Println(err)
	}
	data, err := os.ReadFile(wd + "/staticdata/" + resource_type + "/" + req.Name + "/" + R_list[0].Avatar_url.String)
	if err != nil {
		if err == io.EOF {

		} else {
			log.Fatalln(err)
		}
	}
	var data_base64 string
	data_base64 = base64.StdEncoding.EncodeToString(data)
	resp.Data = data_base64
	return
}

// 先根据用户名查用户id,然后再查询
func GetResource_Signature(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	var R_list []*models.AvatarResource
	var user []*models.User_Total
	prequery := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where u_name=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &user, prequery, req.Name)
	if err != nil {
		log.Println(err)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(user) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find user!"
		return
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find user!"
	}
	query := "select uid,signature,avatar_url from user_information where uid=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, user[0].UID)
	if err != nil {
		log.Println(err)
	}
	//查到用户签名后，将其赋值到回复
	resp.Data = R_list[0].Signature.String
	return
}

// 先根据用户名查用户id,然后再修改user_information表,进行个性签名修改
func PostResource_Signature(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	var R_list []*models.AvatarResource
	var user []*models.User_Total
	prequery := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where u_name=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &user, prequery, req.Name)
	if err != nil {
		log.Println(err)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(user) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find user!"
		return
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find user!"
	}
	query := "update user_information set signature=? where uid=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.Post_data, user[0].UID)
	if err != nil {
		log.Println(err)
	}
	//查到用户签名后，将其赋值到回复
	resp.Data = "signature update success!"
	return
}
func GetResource_Article(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {

	resource_type := "article"
	var R_list []*models.ArticleResource
	query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url from article where Article_ID=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.Name)
	if err != nil {
		log.Println(err)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(R_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find user!"
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find user!"
	}

	data, err := os.ReadFile(wd + "/staticdata/" + resource_type + "/" + R_list[0].Article_url.String)
	if err != nil {
		if err == io.EOF {

		} else {
			log.Fatalln(err)
		}
	}
	resp.Data = string(data)
	return

}

// 获取文章标题
func GetResource_Article_head(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	var R_list []*models.ArticleResource
	query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url from article where Article_ID=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.Name)
	if err != nil {
		log.Println(err)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(R_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find article!"
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find article!"
	}
	resp.Data = R_list[0].Head.String
	return

}

// 获取文章列表
func GetResource_Article_list(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	var R_list []*models.ArticleResource
	query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url from article"
	err := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query)
	if err != nil {
		log.Println(err)
	}
	var length int
	if len(R_list) >= const_values.BIGGEST_ARTICLE_NUM {
		length = const_values.BIGGEST_ARTICLE_NUM
	} else {
		length = len(R_list)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(R_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find article!"
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find article!"
	}
	articleList := make([]types.Article_list_item, length) //创建指定长度的文章集合
	for k, _ := range R_list {                             //拷贝
		articleList[k].Article_ID = R_list[k].Article_ID.String
		articleList[k].Head = R_list[k].Head.String
	}
	resp.ListData = articleList
	return

}
func (r *ResourceLogic) Resource(req *types.ResourceReq) (resp *types.ResourceResponse, err error) {
	res := new(types.ResourceResponse)
	resp = res
	//给绝对路径赋值
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	if req.Resource_type == "avatar" {
		GetResource_Avatar(req, wd, resp) //对头像资源处理
		if err != nil {
			log.Println(err)
		}
	} else if req.Resource_type == "article" {
		GetResource_Article(req, wd, resp) //对文章体处理
		if err != nil {
			log.Println(err)
		}
	} else if req.Resource_type == "signature" {
		GetResource_Signature(req, wd, resp) //对签名资源处理
		if err != nil {
			log.Println(err)
		}
	} else if req.Resource_type == "head" {
		GetResource_Article_head(req, wd, resp) //对文章头处理
		if err != nil {
			log.Println(err)
		}
	} else if req.Resource_type == "article-list" { //获取文章记录列表
		GetResource_Article_list(req, wd, resp) //对文章头处理
		if err != nil {
			log.Println(err)
		}
	}
	return
}

// 处理发送的资源
func (r *ResourceLogic) ResourcePOST(req *types.ResourceReq) (resp *types.ResourceResponse, err error) {
	res := new(types.ResourceResponse)
	resp = res
	//给绝对路径赋值
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if req.Resource_type == "avatar" {

	} else if req.Resource_type == "article" {

	} else if req.Resource_type == "signature" {
		PostResource_Signature(req, wd, resp) //对签名资源处理
		if err != nil {
			log.Println(err)
		}
	} else if req.Resource_type == "head" {

	} else if req.Resource_type == "article-list" { //获取文章记录列表

	}
	return
}
