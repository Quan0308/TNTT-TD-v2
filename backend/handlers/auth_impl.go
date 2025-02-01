package handlers

import (
	"context"
	"github.com/Quan0308/main-api/core/base"
	"github.com/Quan0308/main-api/interfaces"
	"github.com/Quan0308/main-api/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type authHandlerImpl struct {
	db *sqlx.DB
}

func NewAuthHandler(db *sqlx.DB) interfaces.AuthHandler {
	return &authHandlerImpl{
		db: db,
	}
}

func (a *authHandlerImpl) LogIn(ctx context.Context, email string, password string) (uuid.UUID, error) {
	var user models.User
	query := "select * from users where email = ?"

	uow := base.NewUnitOfWork(a.db)

	err := uow.Begin()
	if err != nil {
		uow.RollBack()
		return uuid.Nil, err
	}

	err = uow.GetContext(ctx, &user, query, email)

	if err != nil {
		uow.RollBack()
		return uuid.Nil, err
	}

	uow.Commit()
	return user.Id, nil
}
