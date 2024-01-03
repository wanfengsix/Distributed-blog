package models

import "database/sql"

type Notice struct {
	UID         sql.NullString `db:"UID"`
	U_name      sql.NullString `db:"u_name"`
	Author_ID   sql.NullString `db:"author_ID"`
	Author_name sql.NullString `db:"author_name"`
	Head        sql.NullString `db:"head"`
	Type        sql.NullString `db:"type"`
	Is_read     int            `db:"is_read"`
	Notice_ID   sql.NullString `db:"notice_ID"`
}

func (n *Notice) TableName() string {
	return "notice"
}
