package interfaces

import "github.com/jmoiron/sqlx"

type Container interface {
	Init(db *sqlx.DB)
	GetAuthAPI() AuthAPI
}
