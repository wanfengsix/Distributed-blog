package types

type ResourceReq struct {
	Username      string `json:"username" path:"username"`
	Resource_type string `json:"resource_type" path:"type"`
}
type ResourceResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    string `json:"data" `
}
