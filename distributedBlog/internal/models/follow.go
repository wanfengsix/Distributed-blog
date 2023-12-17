package models

import "database/sql"

type Follow struct {
	Followed_id  sql.NullString `db:"followed_id"`
	Following_id sql.NullString `db:"following_id"`
}

func (f *Follow) TableName() string {
	return "follow"
}
