package search

import (
	"context"
	"distributedBlog/internal/const_values"
	"distributedBlog/internal/models"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

/*
定义资源逻辑体
*/
type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var mysqlDB = sqlx.NewSqlConn("mysql", const_values.MYSQLCONNECTION)
var rds = redis.MustNewRedis(redis.RedisConf{
	Host:        "127.0.0.1:6379",
	Type:        "node",
	Pass:        "",
	Tls:         false,
	NonBlock:    false,
	PingTimeout: time.Second,
})

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
搜索请求处理事务：通过请求关键词，先对redis进行查询，redis缓存没有再对表进行查询，查询有关的资源，将其读入，作为
响应体返回给前端页面
*/

func (s *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	var R_list []*models.ArticleResource
	respo := new(types.SearchResponse)
	resp = respo
	query := "SELECT Article_ID, head, date, UID, likes_nums, comment_nums, article_url FROM article WHERE head LIKE CONCAT('%', ?, '%')"

	//redisquery := strings.ReplaceAll(query, "?", req.Name)
	keyHead := "searchhead:" + req.Name // 设置 Redis 缓存的键名
	keyId := "searchId:" + req.Name
	vHead, err := rds.GetCtx(context.Background(), keyHead)
	vId, err := rds.GetCtx(context.Background(), keyId)
	if err == nil {
		resp.CacheHead = vHead
		resp.CacheId = vId
	}

	if vHead == "" || vId == "" { //没有缓存，那么就查询
		err = mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, req.Name)
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
			//将内容放入缓存
			rds.SetCtx(context.Background(), keyHead, articleList[k].Head)
			rds.SetCtx(context.Background(), keyId, articleList[k].Article_ID)
		}
		resp.ListData = articleList
	}
	return
}