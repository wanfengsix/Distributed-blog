package types

type SearchRequest struct {
	Name string `json:"name" path:"name"`
}
type SearchResponse struct {
	Success   bool
	Message   string              `json:"message"`
	Code      int                 `json:"code"`
	ListData  []Article_list_item `json:"data" `
	CacheHead string              `json:"cachehead"`
	CacheId   string              `json:"cachedataid"`
}
