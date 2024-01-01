package types

type WriteReq struct {
	Data string `json:"data"`
	Uid  string `path:"uid"`
	Head string `json:"head"`
}
type WriteResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
}
