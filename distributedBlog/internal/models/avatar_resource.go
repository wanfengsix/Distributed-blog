package models

import "database/sql"

type AvatarResource struct {
	UID        sql.NullString `json:"uid" db:"uid"`
	Signature  sql.NullString `json:"signature" db:"signature"`
	Avatar_url sql.NullString `json:"avatar_url" db:"avatar_url"`
}
