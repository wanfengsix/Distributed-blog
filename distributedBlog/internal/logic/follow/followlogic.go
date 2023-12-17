package follow

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
定义关注逻辑体
*/
type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var mysqlDB = sqlx.NewSqlConn("mysql", const_values.MYSQLCONNECTION)

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
关注事务:
用户在前端以指定用户名，作者名向后端发送请求，
后端分别对用户表分别查询UID，如果存在关注关系，则取消关注，反之则进行逆操作
*/
func (f *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResponse, err error) {

	res := new(types.FollowResponse)
	resp = res
	//获取u_name,根据请求用户名查询
	var A_list []*models.User_Total //自己用户
	var B_list []*models.User_Total //作者用户
	var UID_self string
	var UID_other string
	query := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where u_name=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &A_list, query, req.Username)
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
		UID_self = A_list[0].UID.String
	}
	query_other := "select uid,u_name,following,followed,article_nums,read_nums,comment_nums,likes_nums,level from user where u_name=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &B_list, query_other, req.U_name)
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
		UID_other = B_list[0].UID.String
	}
	//校验是否已经具备关注关系
	follow_query := "SELECT followed_id ,following_id from follow WHERE followed_id=? AND following_id=?"
	var L_list []*models.Follow
	err1 := mysqlDB.QueryRowsCtx(context.Background(), &L_list, follow_query, UID_other, UID_self)
	if err1 != nil { //即使没有结果报错，也不能返回
		log.Println(err1)
	}
	//如果已经具备关注关系，要执行取消关注操作
	if len(L_list) != 0 {
		CancelFollow(req, resp, UID_self, UID_other)
		resp.Message = "false"
	} else {
		SubmitFollow(req, resp, UID_self, UID_other)
		resp.Message = "true"
	}
	return
}

// 提交关注，先增加用户表自己的关注数，再增加用户表作者的粉丝数，最后增加一条关注记录
func SubmitFollow(req *types.FollowReq, resp *types.FollowResponse, UID_self string, UID_other string) {
	//先改自己关注数
	self_query := "UPDATE user SET following = following+1 WHERE UID=?"
	_, err1 := mysqlDB.ExecCtx(context.Background(), self_query, UID_self)
	if err1 != nil {
		log.Println(err1)
		return
	}
	//再改作者粉丝数
	other_query := "UPDATE user SET followed = followed+1 WHERE UID=?"
	_, err1 = mysqlDB.ExecCtx(context.Background(), other_query, UID_other)
	if err1 != nil {
		log.Println(err1)
		return
	}
	//最后增加关注记录
	last_query := "INSERT INTO follow VALUES(?,?)"
	_, err1 = mysqlDB.ExecCtx(context.Background(), last_query, UID_other, UID_self)
	if err1 != nil {
		log.Println(err1)
		return
	}
}

// 取消关注，先减少用户表自己的关注数，再减少用户表作者的粉丝数，最后删除一条关注记录
func CancelFollow(req *types.FollowReq, resp *types.FollowResponse, UID_self string, UID_other string) {
	//先改自己关注数
	self_query := "UPDATE user SET following = following-1 WHERE UID=?"
	_, err1 := mysqlDB.ExecCtx(context.Background(), self_query, UID_self)
	if err1 != nil {
		log.Println(err1)
		return
	}
	//再改作者粉丝数
	other_query := "UPDATE user SET followed = followed-1 WHERE UID=?"
	_, err1 = mysqlDB.ExecCtx(context.Background(), other_query, UID_other)
	if err1 != nil {
		log.Println(err1)
		return
	}
	//最后增加关注记录
	last_query := "DELETE FROM follow WHERE followed_id=? AND following_id=?"
	_, err1 = mysqlDB.ExecCtx(context.Background(), last_query, UID_other, UID_self)
	if err1 != nil {
		log.Println(err1)
		return
	}
}
