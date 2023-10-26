package models

type BaseModel struct {
	ID []byte `json:"id" gorm:"primarykey"`
}
