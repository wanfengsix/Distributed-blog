package types

type LikesReq struct {
	Article_ID string `json:"Article_ID"`
	U_name     string `json:"u_name"`
}
type LikesResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
}
