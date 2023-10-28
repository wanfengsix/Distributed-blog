package models

import "database/sql"

type BaseModel struct {
	ID sql.NullString `json:"UID" gorm:"primarykey" db:"UID"`
}
