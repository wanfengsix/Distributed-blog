package types

type CommentReq struct {
	Comment_ID string `json:"Comment_ID"`

	Comment_content string `json:"Comment_content"`
	Article_ID      string `json:"Article_ID"`
	UID             string `json:"UID"`
}
type CommentResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
}
