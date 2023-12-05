package auth

import (
	"context"
	"distributedBlog/internal/const_values"
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/models"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

/*
登陆事务：
单个登陆事务类关联一个用户模型，通过访存用户登陆模型数据，来进行登陆事物操作
方法：
先通过SQL语句查询全部字段，再做处理
*/
type Auth struct {
	AuthBody       models.User
	Password_Right bool
	Is_Exist       bool
}

/*
处理有关的auth请求逻辑体，由logx，ctx，svcctx组成。
*/
type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 基于Restul API的登录处理
func (a *AuthLogic) Login(req *types.LoginReq) (resp *types.LoginResponse, err error) {
	newAuth := new(Auth)
	res := new(types.LoginResponse)
	resp = res
	var A_list []*models.User
	query := "select u_name,password,secret_protection1,secret_protection2,secret_protection3,is_Admin from register where u_name=?"
	err = mysqlDB.QueryRowsCtx(context.Background(), &A_list, query, req.Username)
	if err != nil {
		fmt.Println(err)

	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(A_list) == 0 {
		newAuth.Is_Exist = false
		resp.Code = 404
		resp.Success = false
		resp.Message = "can't find user!"
	} else {
		newAuth.AuthBody = *A_list[0]
		newAuth.Is_Exist = true
		resp.Code = 200
		resp.Success = true
		resp.Message = "find user!"
	}

	//校验密码
	if newAuth.Password_IS_Right(req.Password) {
		newAuth.Password_Right = true

	} else {
		resp.Success = false
		newAuth.Password_Right = false
		//json
		resp.Message = "Password Wrong!"
	}

	return
}

var mysqlDB = sqlx.NewSqlConn("mysql", const_values.MYSQLCONNECTION)

func (a *Auth) GetProtectionlist() *[][]byte { //返回用户所有秘保答案列表
	query := "select u_name,password,secret_protection1,secret_protection2,secret_protection3,is_Admin from register"
	err := mysqlDB.QueryRowCtx(context.Background(), &a.AuthBody, query)
	if err != nil {
		panic(err)
	}
	var Protectionlist [][]byte
	Protectionlist = append(Protectionlist, []byte(a.AuthBody.Secret_protection1.String))
	Protectionlist = append(Protectionlist, []byte(a.AuthBody.Secret_protection2.String))
	Protectionlist = append(Protectionlist, []byte(a.AuthBody.Secret_protection3.String))
	return &Protectionlist
}

// 校验密码
func (a *Auth) Password_IS_Right(t string) bool {
	str := base.Hash_encoder(a.AuthBody.Password.String)
	if str == t {
		return true
	} else {
		return false
	}
}

func GetAuth(userName string, passWord string) *Auth { //获取Auth事务体
	newAuth := new(Auth)
	var A_list []*models.User
	query := "select u_name,password,secret_protection1,secret_protection2,secret_protection3,is_Admin from register where u_name=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &A_list, query, userName)
	if err != nil {
		fmt.Println(err)

	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(A_list) == 0 {
		newAuth.Is_Exist = false
	} else {
		newAuth.AuthBody = *A_list[0]
		newAuth.Is_Exist = true
	}

	//校验密码
	if newAuth.Password_IS_Right(passWord) {
		newAuth.Password_Right = true

	} else {
		newAuth.Password_Right = false
		//json
	}
	return newAuth

}
func (a *Auth) isEqual(l Auth, r Auth) bool { //比对两个auth请求
	if l.AuthBody.Is_Admin != r.AuthBody.Is_Admin {
		return false
	}
	if l.AuthBody.Password != r.AuthBody.Password {
		return false
	}
	if l.AuthBody.Secret_protection1 != r.AuthBody.Secret_protection1 {
		return false
	}
	if l.AuthBody.Secret_protection2 != r.AuthBody.Secret_protection2 {
		return false
	}
	if l.AuthBody.Secret_protection3 != r.AuthBody.Secret_protection3 {
		return false
	}
	return true
}
func (a *Auth) AuthLogic(req *types.LoginReq) (resp *types.Response, err error) {
	another_a := a.GetProtectionlist()
	fmt.Printf("前端用户名：%s\n", req.Username)
	fmt.Printf("前端密码：%s\n", req.Password)
	if another_a == nil {
		log.Fatalln("数据库查询错误！")
		return
	}
	fmt.Println("确认登陆！密保如下：")
	fmt.Println(string((*another_a)[0]))
	fmt.Println(string((*another_a)[1]))
	fmt.Println(string((*another_a)[2]))
	return
}
