package models

import "database/sql"

type User_Total struct {
	UID          sql.NullString `json:"uid" db:"uid"`
	U_name       sql.NullString `json:"u_name" db:"u_name"`
	Following    int            `json:"following" db:"following"`
	Followed     int            `json:"followed" db:"followed"`
	Article_nums int            `json:"article_nums" db:"article_nums"`
	Read_nums    int            `json:"read_nums" db:"read_nums"`
	Comment_nums int            `json:"comment_nums" db:"comment_nums"`
	Likes_nums   int            `json:"likes_nums" db:"likes_nums"`
	Level        int            `json:"level" db:"level"`
}

func (u *User_Total) TableName() string {
	return "user"
}
