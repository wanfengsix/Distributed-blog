package types

type LikesReq struct {
	U_name     string `json:"u_name"`
	Article_ID string `json:"Article_ID"`
}
type LikesResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
}
