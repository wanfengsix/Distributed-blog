package types

type WriteReq struct {
	Data      string `json:"data"`
	Uid       string `path:"uid"`
	Head      string `json:"head"`
	ArticleId string `json:"article_id"`
}
type WriteResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
}
