package login

import (
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/logic/auth"
	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
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

// 基于restful api的登录处理
func Rest_LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理跨域请求
		base.Prefix_Managing(w, r)
		if r.Method != "POST" {
			log.Fatalln("请求类型错误！请检查路由配置")
			return
		}
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		a := auth.NewAuthLogic(r.Context(), svcCtx)
		resp, err := a.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
		/*
			if resp.Code == 200 {
				//发送正确到前端
				base.GetDataHandler(&w, r, &types.LoginResponse{Success: true, Message: "Password is correct!"})
			} else if resp.Message == "can't find user!" {
				//发送错误到前端
				base.GetDataHandler(&w, r, &types.LoginResponse{Success: false, Message: "User is not exist!", Code: 1001})
			} else if resp.Message == "Password Wrong!" {
				//发送错误到前端
				base.GetDataHandler(&w, r, &types.LoginResponse{Success: false, Message: "Password is incorrect!", Code: 1001})
			}
		*/
	}
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
			base.GetDataHandler(&w, r, &types.LoginResponse{Success: true, Message: "Password is correct!"})
		} else if !a.Is_Exist {
			//发送错误到前端
			base.GetDataHandler(&w, r, &types.LoginResponse{Success: false, Message: "User is not exist!", Code: 1001})
		} else if !a.Password_Right {
			//发送错误到前端
			base.GetDataHandler(&w, r, &types.LoginResponse{Success: false, Message: "Password is incorrect!", Code: 1001})
		}
	}
}
