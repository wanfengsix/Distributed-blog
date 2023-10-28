package auth

import (
	"context"
	"distributedBlog/internal/models"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

/*
登陆事务：
单个登陆事务类关联一个用户模型，通过访存用户登陆模型数据，来进行登陆事物操作
方法：
先通过SQL语句查询全部字段，再做处理
*/
type Auth struct {
	AuthBody models.User
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
