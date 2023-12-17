package types

type ResourceReq struct {
	Name          string `json:"name" path:"name"`
	Resource_type string `json:"resource_type" path:"type"`
	Post_data     string `json:"post_data"` //仅在POST请求用到
}

// 用来返回的格式 这些item归根到底还是response
type Article_list_item struct {
	Head       string `json:"head"`
	Article_ID string `json:"article_id"`
}
type Comment_list_item struct {
	Comment_ID      string `json:"Comment_ID"`
	Comment_content string `json:"Comment_content"`
	UID             string `json:"UID"`
	Date            string `json:"date"`
}
type ResourceResponse struct {
	Success         bool
	Message         string              `json:"message"`
	Code            int                 `json:"code"`
	Data            string              `json:"data" `
	ArticleListData []Article_list_item `json:"article_list"` //只有请求文章列表集合用到
	CommentListData []Comment_list_item `json:"comment_list"` //只有请求评论列表集合用到
}
