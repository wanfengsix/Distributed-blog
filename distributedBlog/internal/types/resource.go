package types

type ResourceReq struct {
	Name          string `json:"name" path:"name"`
	Resource_type string `json:"resource_type" path:"type"`
	Post_data     string `json:"post_data"` //仅在POST请求用到
}
type Article_list_item struct {
	Head       string `json:"head"`
	Article_ID string `json:"article_id"`
}
type ResourceResponse struct {
	Success  bool
	Message  string              `json:"message"`
	Code     int                 `json:"code"`
	Data     string              `json:"data" `
	ListData []Article_list_item `json:"article_list"` //只有请求列表集合用到
}
