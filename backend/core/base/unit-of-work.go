package base

import (
	"context"
	"database/sql"
	"log"

	"github.com/Quan0308/main-api/interfaces"
	"github.com/jmoiron/sqlx"
)

type UnitOfWorkImpl struct {
	Db *sqlx.DB
	Tx *sqlx.Tx
}

func NewUnitOfWorkImpl(db *sqlx.DB) interfaces.UnitOfWork {
	return &UnitOfWorkImpl{Db: db}
}

func (u *UnitOfWorkImpl) Begin() error {
	tx, err := u.Db.Beginx()
	if err != nil {
		log.Printf("Error begin")
		u.RollBack()
		return err
	}

	u.Tx = tx
	return nil
}

func (u *UnitOfWorkImpl) Commit() error {
	err := u.Tx.Commit()
	if err != nil {
		u.RollBack()
		log.Printf("Error Commit")
		return err
	}
	return nil
}

func (u *UnitOfWorkImpl) RollBack() error {
	err := u.Tx.Rollback()
	if err != nil {
		log.Printf("Error RollBack")
		return err
	}
	return nil
}

func (u *UnitOfWorkImpl) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return u.Tx.ExecContext(ctx, query, args...)
}

func (u *UnitOfWorkImpl) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return u.Tx.GetContext(ctx, dest, query, args...)
}

func (u *UnitOfWorkImpl) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return u.Tx.SelectContext(ctx, dest, query, args...)
}
