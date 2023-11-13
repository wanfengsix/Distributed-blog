package login

import (
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/logic/auth"
	"distributedBlog/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func loginHandler(authReq *types.LoginTokenInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//处理跨域请求
		base.CORS_Managing(&w, r)
		fmt.Fprintf(w, "Password is correct!")
		//base.CORS_ALLOW_ALL(w, r)
	}
}
func Prefix_login_managing() http.HandlerFunc {
	return base.Prefix_Managing
}

// 处理post请求：先进行跨域请求处理，之后再通过io对reuqest的body进行解析出json数据，从而获取账户和密码
func PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "POST" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		//读取并按json进行解析
		body, err2 := io.ReadAll(r.Body)
		if err2 != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		var loginData types.LoginReq             // 定义用于存储解析后的JSON数据的结构体
		err3 := json.Unmarshal(body, &loginData) // 解析JSON数据
		if err3 != nil {
			http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
			return
		}
		username := loginData.Username // 获取解析后的用户名
		password := loginData.Password // 获取解析后的密码
		fmt.Println("用户名:" + username)
		fmt.Println("密码:" + password)
		//访问数据库进行校验，开始数据库事务
		var a auth.Auth
		a = *auth.GetAuth(username, password)
		fmt.Println(a.Password_Right)
		if a.Password_Right {
			//发送正确到前端
			base.GetDataHandler(&w, r, "Password is correct!")
		} else {
			//发送错误到前端
			base.GetDataHandler(&w, r, "Password is incorrect!")
		}
	}
}
