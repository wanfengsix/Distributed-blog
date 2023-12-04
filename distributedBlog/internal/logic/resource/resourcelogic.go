package resource

import (
	"context"
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

var mysqlDB = sqlx.NewSqlConn("mysql", "root:xin365118@tcp(127.0.0.1:3306)/dusha?charset=utf8mb4&parseTime=True&loc=Local")

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
func (r *ResourceLogic) Resource(req *types.ResourceReq) (resp *types.ResourceResponse, err error) {
	res := new(types.ResourceResponse)
	resp = res
	var R_list []*models.Resource
	resource_type := ""
	if req.Resource_type == "avatar" {
		resource_type = "avatar"
	} else if req.Resource_type == "article" {
		resource_type = "article"
	}
	query := "select uid,signature,avatar_url from user_information where uid=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.Username)
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
	//给绝对路径赋值
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	data, err := os.ReadFile(wd + "/staticdata/" + req.Username + "/" + resource_type + "/" + R_list[0].Avatar_url.String)
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
