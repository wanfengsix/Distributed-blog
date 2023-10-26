package models

type User struct {
	BaseModel
	password           []byte `json:"password"`
	secret_protection1 []byte `json:"secret_protection1"`
	secret_protection2 []byte `json:"secret_protection2"`
	secret_protection3 []byte `json:"secret_protection3"`
	is_Admin           int    `json:"is_Admin"`
}

func (User) TableName() string {
	return "user"
}
