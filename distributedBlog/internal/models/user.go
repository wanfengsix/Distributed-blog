package models

import (
	"database/sql"
)

type User struct {
	UID                sql.NullString `json:"uid" db:"uid"`
	Password           sql.NullString `json:"password" db:"password"`
	Secret_protection1 sql.NullString `json:"secret_protection1" db:"secret_protection1"`
	Secret_protection2 sql.NullString `json:"secret_protection2" db:"secret_protection2"`
	Secret_protection3 sql.NullString `json:"secret_protection3" db:"secret_protection3"`
	Is_Admin           int64          `json:"is_Admin" db:"is_Admin"`
}

func (u *User) TableName() string {
	return "user"
}
