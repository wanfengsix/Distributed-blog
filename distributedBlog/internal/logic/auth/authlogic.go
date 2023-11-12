package auth

import (
	"context"
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/models"
	"distributedBlog/internal/types"
	"fmt"
	"log"

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
}

var mysqlDB = sqlx.NewSqlConn("mysql", "root:xin365118@tcp(127.0.0.1:3306)/dusha?charset=utf8mb4&parseTime=True&loc=Local")

func (a *Auth) GetProtectionlist() *[][]byte { //返回用户所有秘保答案列表
	query := "select uid,password,secret_protection1,secret_protection2,secret_protection3,is_Admin from register"
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
	query := "select uid,password,secret_protection1,secret_protection2,secret_protection3,is_Admin from register where uid=?"
	err := mysqlDB.QueryRowCtx(context.Background(), &newAuth.AuthBody, query, userName)
	if err != nil {
		panic(err)
	}
	//校验密码
	if newAuth.Password_IS_Right(passWord) {
		newAuth.Password_Right = true
	} else {
		newAuth.Password_Right = false
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
