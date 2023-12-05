package types

type ResourceReq struct {
	Name          string `json:"name" path:"name"`
	Resource_type string `json:"resource_type" path:"type"`
}
type ResourceResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    string `json:"data" `
}
