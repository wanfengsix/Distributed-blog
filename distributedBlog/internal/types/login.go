package types

type LoginTokenInfo struct {
	GranterType string `form:"granter_type"`
	LoginReq
}
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
}
