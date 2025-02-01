package handlers

import (
	"github.com/Quan0308/main-api/interfaces"
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

func (a *authHandlerImpl) SignUp(email string, password string) (string, error) {
	return "success", nil
}

func (a *authHandlerImpl) SignIn(email string, password string) (uuid.UUID, error) {
	query := "select id from users where email = ?"
	var userId uuid.UUID
	err := a.db.QueryRow(query, email).Scan(&userId)

	if err != nil {
		return uuid.Nil, err
	}
	return userId, nil
}
