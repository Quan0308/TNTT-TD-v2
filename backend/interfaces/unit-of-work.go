package interfaces

import (
	"context"
	"database/sql"
)

type UnitOfWork interface {
	Begin() error
	Commit() error
	RollBack() error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
