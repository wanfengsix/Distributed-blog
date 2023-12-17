package models

import "database/sql"

type Comment struct {
	Comment_ID      sql.NullString `db:"Comment_ID"`
	Comment_content sql.NullString `db:"Comment_content"`
	Article_ID      sql.NullString `db:"Article_ID"`
	UID             sql.NullString `db:"UID"`
	Date            sql.NullString `db:"date"`
}

func (c *Comment) TableName() string {
	return "comment"
}
