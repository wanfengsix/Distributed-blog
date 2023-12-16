package models

import "database/sql"

//用来访问数据库
type CommentResource struct {
	Comment_ID      sql.NullString `json:"Comment_ID" db:"Comment_ID"`
	Comment_content sql.NullString `json:"Comment_content" db:"Comment_content"`
	Article_ID      sql.NullString `json:"Article_ID" db:"Article_ID"`
	UID             sql.NullString `json:"UID" db:"UID"`
}

func (a *CommentResource) TableName() string {
	return "comment"
}
