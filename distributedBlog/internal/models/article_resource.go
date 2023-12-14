package models

import "database/sql"

type ArticleResource struct {
	Article_ID   sql.NullString `json:"Article_ID" db:"Article_ID"`
	Head         sql.NullString `json:"head" db:"head"`
	Date         sql.NullString `json:"date" db:"date"`
	UID          sql.NullString `json:"UID" db:"UID"`
	Likes_nums   int            `json:"likes_nums" db:"likes_nums"`
	Comment_nums int            `json:"comment_nums" db:"comment_nums"`
	Article_url  sql.NullString `json:"article_url" db:"article_url"`
}

func (a *ArticleResource) TableName() string {
	return "article"
}
