package types

import "distributedBlog/internal/models"

type UserReq struct {
	Type  string `path:"type"`
	Value string `path:"value"`
}
type UserResponse struct {
	Success bool
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Data    models.User_Total `json:"data" ` //用户体
}
