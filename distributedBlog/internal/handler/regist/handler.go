package regist

import (
	"distributedBlog/internal/handler/base"
	"distributedBlog/internal/logic/regist"
	"distributedBlog/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Prefix_regist_managing() http.HandlerFunc {
	return base.Prefix_Managing
}

// 处理post请求：先进行跨域请求处理，之后再通过io对reuqest的body进行解析出json数据，从而获取注册信息,若无该用户名，进而将信息插入到数据库中
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
		var RegistData types.RegistReq            // 定义用于存储解析后的JSON数据的结构体
		err3 := json.Unmarshal(body, &RegistData) // 解析JSON数据
		if err3 != nil {
			http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
			return
		}
		username := RegistData.Username                         // 获取解析后的用户名
		password := RegistData.Password                         // 获取解析后的密码
		password_protection1 := RegistData.Password_Protection1 //获取解析后的密保1
		password_protection2 := RegistData.Password_Protection2 //获取解析后的密保2
		password_protection3 := RegistData.Password_Protection3 //获取解析后的密保3
		fmt.Println("用户名:" + username)
		fmt.Println("密码:" + password)
		fmt.Println("密保1:" + password_protection1)
		fmt.Println("密保1:" + password_protection2)
		fmt.Println("密保1:" + password_protection3)
		//访问数据库进行校验，开始数据库事务
		var re regist.Regist
		re = *regist.GetRegist(RegistData)
		fmt.Println(re.Is_Exist)
		if !re.Is_Exist {
			//发送正确到前端
			base.GetDataHandler(&w, r, &types.LoginResponse{Success: true, Message: "Regist is success!"})
		} else {
			//发送错误到前端
			base.GetDataHandler(&w, r, &types.LoginResponse{Success: false, Message: "User is exist!", Code: 1001})
		}
	}
}
