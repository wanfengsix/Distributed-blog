package regist

import (
	"context"
	"distributedBlog/internal/models"
	"distributedBlog/internal/types"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

/*
注册事务：
单个登陆事务类关联一个用户模型，通过访存用户模型数据，来进行注册事物操作
方法：
先通过SQL语句查询全部字段，再做处理
*/
type Regist struct {
	RegistBody models.User
	Is_Exist   bool
}

var mysqlDB = sqlx.NewSqlConn("mysql", "root:xin365118@tcp(127.0.0.1:3306)/dusha?charset=utf8mb4&parseTime=True&loc=Local")

func GetRegist(R types.RegistReq) *Regist { //获取Regist事务体
	newRegist := new(Regist)
	var R_list []*models.User
	query := "select u_name,password,secret_protection1,secret_protection2,secret_protection3,is_Admin from register where u_name=?"
	err := mysqlDB.QueryRowsCtx(context.Background(), &R_list, query, R.Username)
	if err != nil {
		fmt.Println(err)

	}
	//检验会不会查询的到,如果查询不到，那么就返回不存在
	if len(R_list) == 0 {
		newRegist.Is_Exist = false
	} else {
		newRegist.RegistBody = *R_list[0]
		newRegist.Is_Exist = true
	}

	//如果不存在就插入一条用户注册数据，否则什么也不做
	if !newRegist.Is_Exist {
		mysqlDB.ExecCtx(context.Background(), "INSERT INTO register VALUES (?,?,?,?,?,0)", R.Username, R.Password, R.Password_Protection1, R.Password_Protection2, R.Password_Protection3)
	}
	return newRegist

}
