package models

import "database/sql"

type Draft struct {
	UID        sql.NullString `db:"UID"`
	Date       sql.NullString `db:"draft_date"`
	Draft_Head sql.NullString `db:"draft_head"`
	Draft_url  sql.NullString `db:"draft_url"`
}
