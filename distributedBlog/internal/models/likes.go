package models

import "database/sql"

type Likes struct {
	UID        sql.NullString ` db:"uid"`
	Article_ID sql.NullString `db:"Article_ID"`
}

func (l *Likes) TableName() string {
	return "likes"
}
