package types

type FollowReq struct {
	U_name   string `json:"u_name"`
	Username string `json:"username"`
}

type FollowResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
}
