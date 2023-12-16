package types

import "distributedBlog/internal/models"

type UserReq struct {
	Type  string `path:"type"`
	Value string `path:"value"`
}
type UserList_Item struct {
	UID    string `json:"UID" db:"uid"`
	U_name string `json:"u_name" db:"u_name"`
}
type UserResponse struct {
	Success bool
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Data    models.User_Total `json:"data" ` //用户体
	List    []UserList_Item   `json:"List"`  //仅在请求用户列表时用到
}
