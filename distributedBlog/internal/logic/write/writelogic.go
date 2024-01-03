package write

import (
	"context"
	"distributedBlog/internal/const_values"
	"distributedBlog/internal/models"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

/*
定义撰写逻辑体
*/
type WriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var mysqlDB = sqlx.NewSqlConn("mysql", const_values.MYSQLCONNECTION)

func NewWriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WriteLogic {
	return &WriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
搜索请求处理事务：通过uid，对article表进行插入，根据请求的文章内容通过article+随机数的方式保存文件，将文件存储结果作为
响应体返回给前端页面
*/
func (w *WriteLogic) Write(req *types.WriteReq) (resp *types.WriteResponse, err error) {
	respo := new(types.WriteResponse)
	resp = respo

	if req.ArticleId != "" { //不为空，则是对现有文章进行修改，查询该文章，对其标题内容修改
		var R_list []*models.ArticleResource
		query := "select Article_ID,head,date,UID,likes_nums,comment_nums,article_url,abstract,is_visible from article where Article_ID=?"
		err := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.ArticleId)
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
		//写入文件内容
		//给绝对路径赋值
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		err = os.WriteFile(wd+"/staticdata/"+"article"+"/"+R_list[0].Article_url.String, []byte(req.Data), os.ModePerm)
		if err != nil {
			if err == io.EOF {

			} else {
				log.Println(err)
			}
		}
		//修改文章标题
		update := "UPDATE article SET head=? WHERE Article_ID=?"
		_, err = mysqlDB.ExecCtx(context.Background(), update, req.Head, req.ArticleId)
		if err != nil {
			log.Println(err)
		}
	} else {
		var U_list []*models.User_Total
		user_query := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where UID=?"
		err2 := mysqlDB.QueryRowsCtx(context.Background(), &U_list, user_query, req.Uid)
		var UserName string
		if err2 != nil {
			log.Println(err2)
		}
		//检验会不会查询的到,如果查询不到，那么就返回不存在
		if len(U_list) == 0 {
			resp.Code = 404
			resp.Success = false
			resp.Message = "can't find user!"
			return
		} else {
			resp.Code = 200
			resp.Success = true
			resp.Message = "find user!"
			UserName = U_list[0].UID.String
		}
		NewArticleId := UserName + strconv.Itoa(rand.Intn(100))
		mysqlDateFormat := "2006-01-02"
		time_now := time.Now().Format(mysqlDateFormat)
		FileName := NewArticleId + ".txt" //生成随机文件名
		//文章摘要
		abstract := req.Data[:50]
		//插入文章记录
		query := "INSERT INTO article VALUES(?,?,?,?,?,?,?,?,?)"
		_, err2 = mysqlDB.ExecCtx(context.Background(), query, NewArticleId, req.Head, time_now, req.Uid, 0, 0, FileName, abstract, 0)
		if err2 != nil {
			log.Println(err2)
		}
		//写入文件内容
		//给绝对路径赋值
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		wd = wd + "/staticdata" + "/article"
		RealPath := filepath.Join(wd, FileName)
		err = os.WriteFile(RealPath, []byte(req.Data), os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	return
}
func (w *WriteLogic) Draft(req *types.WriteReq) (resp *types.WriteResponse, err error) {
	respo := new(types.WriteResponse)
	resp = respo
	var U_list []*models.User_Total
	user_query := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where UID=?"
	err2 := mysqlDB.QueryRowsCtx(context.Background(), &U_list, user_query, req.Uid)
	var UserName string
	if err2 != nil {
		log.Println(err2)
	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(U_list) == 0 {
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find user!"
		return
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find user!"
		UserName = U_list[0].UID.String
	}

	//先查询草稿记录有没有
	var R_list []*models.Draft
	query := "select UID,draft_date,draft_head,draft_url from draft where UID=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.Uid)
	if err != nil {
		log.Println(err)
	}
	//检验会不会查询的到,如果查询到，那么就执行更新操作
	if len(R_list) == 0 {
		resp.Code = 200
		resp.Success = true
		resp.Message = "new draft!"
		NewArticleId := UserName + strconv.Itoa(rand.Intn(100))
		mysqlDateFormat := "2006-01-02"
		time_now := time.Now().Format(mysqlDateFormat)
		FileName := NewArticleId + ".txt" //生成随机文件名
		//插入草稿记录
		query := "INSERT INTO draft VALUES(?,?,?,?)"
		_, err2 = mysqlDB.ExecCtx(context.Background(), query, req.Uid, time_now, req.Head, FileName)
		if err2 != nil {
			log.Println(err2)
			//终止写入文件内容
			return
		}
		//写入文件内容
		//给绝对路径赋值
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		wd = wd + "/staticdata" + "/draft"
		RealPath := filepath.Join(wd, FileName)
		err = os.WriteFile(RealPath, []byte(req.Data), os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	} else {
		resp.Code = 200
		resp.Success = true
		resp.Message = "find draft!"
		//写入文件内容
		//给绝对路径赋值
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		wd = wd + "/staticdata" + "/draft"
		RealPath := filepath.Join(wd, R_list[0].Draft_url.String)
		err = os.WriteFile(RealPath, []byte(req.Data), os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	return
}

// 获取草稿内容
func (w *WriteLogic) GetResource_Draft(req *types.WriteReq) (resp *types.WriteResponse, err error) {
	respo := new(types.WriteResponse)
	resp = respo
	resource_type := "draft"
	var R_list []*models.Draft
	query := "select UID,draft_date,draft_head,draft_url from draft where UID=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.Uid)
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
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	data, err := os.ReadFile(wd + "/staticdata/" + resource_type + "/" + R_list[0].Draft_url.String)
	if err != nil {
		if err == io.EOF {

		} else {
			log.Fatalln(err)
		}
	}
	resp.Message = string(data)
	return
}
