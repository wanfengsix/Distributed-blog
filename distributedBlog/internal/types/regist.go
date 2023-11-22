package types

type RegistReq struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	Password_Protection1 string `json:password_protection1`
	Password_Protection2 string `json:password_protection2`
	Password_Protection3 string `json:password_protection3`
}
type RegistResponse struct {
	Success bool
	Message string `json:"message"`
	Code    int    `json:"code"`
}
