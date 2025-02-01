package base

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type UnitOfWork struct {
	Db *sqlx.DB
	Tx *sqlx.Tx
}

func NewUnitOfWork(db *sqlx.DB) *UnitOfWork {
	return &UnitOfWork{Db: db}
}

func (u *UnitOfWork) Begin() error {
	tx, err := u.Db.Beginx()
	if err != nil {
		log.Printf("Error begin")
		return err
	}

	u.Tx = tx
	return nil
}

func (u *UnitOfWork) Commit() error {
	err := u.Tx.Commit()
	if err != nil {
		log.Printf("Error Commit")
		return err
	}
	return nil
}

func (u *UnitOfWork) RollBack() error {
	err := u.Tx.Rollback()
	if err != nil {
		log.Printf("Error RollBack")
		return err
	}
	return nil
}

func (u *UnitOfWork) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return u.Tx.ExecContext(ctx, query, args...)
}

func (u *UnitOfWork) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return u.Tx.GetContext(ctx, dest, query, args...)
}

func (u *UnitOfWork) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return u.Tx.SelectContext(ctx, dest, query, args...)
}
