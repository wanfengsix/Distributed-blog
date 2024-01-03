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
	"strconv"

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

// 获取文章内容
func GetResource_Article(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {

	resource_type := "article"
	var R_list []*models.ArticleResource
	query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url,abstract,is_visible from article where Article_ID=?"
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

	data, err := os.ReadFile(wd + "/staticdata/" + resource_type + "/" + R_list[0].Article_url.String)
	if err != nil {
		if err == io.EOF {

		} else {
			log.Println(err)
		}
	}
	resp.Data = string(data)
	return

}

// 获取文章标题
func GetResource_Article_head(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	var R_list []*models.ArticleResource
	query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url,abstract,is_visible from article where Article_ID=?"
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
	query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url,abstract,is_visible from article"
	err := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query)
	if err != nil {
		log.Println(err)
	}
	var length int
	if len(R_list) >= const_values.BIGGEST_LIST_NUM {
		length = const_values.BIGGEST_LIST_NUM
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
	for k := 0; k < length; k++ {                          //拷贝
		articleList[k].Article_ID = R_list[k].Article_ID.String
		articleList[k].Head = R_list[k].Head.String
	}
	resp.ArticleListData = articleList
	return

}

// 获取评论列表
func GetResource_Comment_list(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	var R_list []*models.Comment
	query := "select Comment_ID,Comment_content,Article_ID,UID,date from comment where Article_ID=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.Name)
	if err != nil {
		log.Println(err)
	}
	var length int
	if len(R_list) >= const_values.BIGGEST_LIST_NUM {
		length = const_values.BIGGEST_LIST_NUM
	} else {
		length = len(R_list)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(R_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find Comment!"
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find Comment!"
	}
	commentList := make([]types.Comment_list_item, length) //创建指定长度的文章集合
	for k := 0; k < length; k++ {                          //拷贝
		commentList[k].Comment_ID = R_list[k].Comment_ID.String
		commentList[k].Comment_content = R_list[k].Comment_content.String
		commentList[k].UID = R_list[k].UID.String
		commentList[k].Date = R_list[k].Date.String
	}
	resp.CommentListData = commentList
	return
}

// 获取文章点赞数，通过查询article表获取
func GetResource_likes_nums_article(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	var R_list []*models.ArticleResource
	query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url ,abstract,is_visible from article where Article_ID=?"
	err1 := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.Name)
	if err1 != nil {
		log.Println(err1)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(R_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find Article Or Likes!"
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find Article Or Likes!"
	}
	resp.Data = strconv.Itoa(R_list[0].Likes_nums) //转换为字符串
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
	} else if req.Resource_type == "article" {
		GetResource_Article(req, wd, resp) //对文章体处理
	} else if req.Resource_type == "signature" {
		GetResource_Signature(req, wd, resp) //对签名资源处理
	} else if req.Resource_type == "head" {
		GetResource_Article_head(req, wd, resp) //对文章头处理
	} else if req.Resource_type == "article-list" { //获取文章记录列表
		GetResource_Article_list(req, wd, resp) //对文章列表处理
	} else if req.Resource_type == "comment-list" { //获取评论列表
		GetResource_Comment_list(req, wd, resp) //对评论列表处理
	} else if req.Resource_type == "likes-nums-article" { //获取文章的点赞数
		GetResource_likes_nums_article(req, wd, resp) //对文章的点赞数处理
	} else if req.Resource_type == "author-rank" { //作者榜
		GetResource_Author_Rank(req, resp)
	} else if req.Resource_type == "article-rank" { //作者榜
		GetResource_Article_Rank(req, resp)

	} else if req.Resource_type == "delete-article" { //删除文章
		Delete_Article(req, resp)
	}
	return
}

// 根据article_id删除文章:先查article表查找文章文件位置，再删除文件，最后删除记录
func Delete_Article(req *types.ResourceReq, resp *types.ResourceResponse) {
	//先删除文件
	var Article_list []*models.ArticleResource
	//查询文章记录，获取url
	get_query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url,abstract,is_visible from article where Article_ID=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &Article_list, get_query, req.Name)
	if err != nil {
		log.Println(err)
	}
	var url string
	if len(Article_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find article!"
		return
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find article!"
		url = Article_list[0].Article_url.String
	}
	//执行删除文件操作
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	wd = wd + "/staticdata" + "/article/" + url
	err = os.Remove(wd)

	//对应的用户点赞数-1,减去相应的评论数
	query_4 := "select UID ,Article_ID from likes where Article_ID=?  "
	var User_list []*models.User_Total
	err = mysqlDB.QueryRowsCtx(context.Background(), &User_list, query_4, req.Name)
	if err != nil {
		log.Println(err)
	}
	for _, v := range User_list {
		likessub_query := "update user set likes_nums=likes_nums-1 where UID=?"
		_, err = mysqlDB.ExecCtx(context.Background(), likessub_query, v.UID.String)
		if err != nil {
			log.Println(err)
		}
		//搜集各个用户在这篇文章评论数
		var comment_list []*models.Comment
		commentSearch_query := "select comment_ID,comment_content,Article_ID,UID,date from comment where UID=?"
		err = mysqlDB.QueryRowsCtx(context.Background(), &comment_list, commentSearch_query, v.UID.String)
		if err != nil {
			log.Println(err)
		}
		//减去相应评论数
		commentsub_query := "update user set comment_nums= comment_nums-? where UID=?"
		_, err = mysqlDB.ExecCtx(context.Background(), commentsub_query, len(comment_list), v.UID.String)
		if err != nil {
			log.Println(err)
		}
	}

	//删除点赞和评论记录
	query2 := "delete from comment where Article_ID=?"
	_, err = mysqlDB.ExecCtx(context.Background(), query2, req.Name)
	if err != nil {
		log.Println(err)
	}
	query3 := "delete from likes where Article_ID=?"
	_, err = mysqlDB.ExecCtx(context.Background(), query3, req.Name)
	if err != nil {
		log.Println(err)
	}
	//删除文章记录
	query := "delete from article where Article_ID=?"
	_, err = mysqlDB.ExecCtx(context.Background(), query, req.Name)
	if err != nil {
		log.Println(err)
	}
}

// 获取作者榜，根据粉丝数来递减查询，取前8位
func GetResource_Author_Rank(req *types.ResourceReq, resp *types.ResourceResponse) {
	var R_list []*models.User_Total
	query := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user order by followed DESC"
	err := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query)
	if err != nil {
		log.Println(err)
	}
	var length int
	if len(R_list) >= const_values.BIGGEST_LIST_NUM {
		length = const_values.BIGGEST_LIST_NUM
	} else {
		length = len(R_list)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(R_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find author!"
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find author!"
	}
	authorList := make([]types.Author_list_item, length) //创建指定长度的文章集合
	for k := 0; k < length; k++ {                        //拷贝
		authorList[k].Name = R_list[k].U_name.String
		authorList[k].UID = R_list[k].UID.String
	}
	resp.AuthorListData = authorList
	return
}

// 获取作者榜，根据粉丝数来递减查询，取前8位
func GetResource_Article_Rank(req *types.ResourceReq, resp *types.ResourceResponse) {
	var R_list []*models.ArticleResource
	query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url,abstract,is_visible from article order by likes_nums DESC"
	err := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query)
	if err != nil {
		log.Println(err)
	}
	var length int
	if len(R_list) >= const_values.BIGGEST_LIST_NUM {
		length = const_values.BIGGEST_LIST_NUM
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
	for k := 0; k < length; k++ {                          //拷贝
		articleList[k].Article_ID = R_list[k].Article_ID.String
		articleList[k].Head = R_list[k].Head.String
	}
	resp.ArticleListData = articleList
	return

}

// 获取用户对文章的点赞状态，先通过用户名查UID，通过查询likes表获取
func GetResource_isliked(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	//获取u_name,根据请求用户名查询
	var A_list []*models.User_Total
	var UID string
	query := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where u_name=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &A_list, query, req.Name)
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
	//校验是否已经具备点赞关系
	liked_query := "SELECT uid ,Article_ID from likes WHERE UID=? AND Article_ID=?"
	var L_list []*models.Likes
	err1 := mysqlDB.QueryRowsCtx(context.Background(), &L_list, liked_query, UID, req.Post_data)
	if err1 != nil { //即使没有结果报错，也不能返回
		log.Println(err1)
	}
	//如果已经具备点赞关系，要对信息赋值
	if len(L_list) != 0 {
		resp.Message = "true"
	} else {
		resp.Message = "false"
	}
	return
}

func GetResource_isFollowed(req *types.ResourceReq, wd string, resp *types.ResourceResponse) {
	//获取u_name,根据请求用户名查询
	var A_list []*models.User_Total
	var UID string
	query := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where u_name=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &A_list, query, req.Name)
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
	//校验是否已经具备关注关系
	liked_query := "SELECT followed_id,following_id from follow WHERE followed_id=? AND following_id=?"
	var L_list []*models.Likes
	err1 := mysqlDB.QueryRowsCtx(context.Background(), &L_list, liked_query, req.Post_data, UID)
	if err1 != nil { //即使没有结果报错，也不能返回
		log.Println(err1)
	}
	//如果已经具备关注关系，要对信息赋值
	if len(L_list) != 0 {
		resp.Message = "true"
	} else {
		resp.Message = "false"
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
	} else if req.Resource_type == "head" {

	} else if req.Resource_type == "article-list" { //获取文章记录列表

	} else if req.Resource_type == "comment-list" { //获取文章评论列表

	} else if req.Resource_type == "islike" { //获取用户对文章的点赞状态
		GetResource_isliked(req, wd, resp) //对文章的点赞数处理
	} else if req.Resource_type == "isFollow" { //获取用户对文章的点赞状态
		GetResource_isFollowed(req, wd, resp) //对文章的点赞数处理
	}
	return
}
