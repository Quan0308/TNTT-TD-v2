package handlers

import (
	"context"

	"github.com/Quan0308/main-api/interfaces"
	"github.com/google/uuid"
)

type authHandlerImpl struct {
	uow      interfaces.UnitOfWork
	userRepo interfaces.UserRepository
}

func NewAuthHandler(uow interfaces.UnitOfWork, userRepo interfaces.UserRepository) interfaces.AuthHandler {
	return &authHandlerImpl{
		uow:      uow,
		userRepo: userRepo,
	}
}

func (a *authHandlerImpl) LogIn(ctx context.Context, email string, password string) (uuid.UUID, error) {
	err := a.uow.Begin()
	if err != nil {
		return uuid.Nil, err
	}

	user, err := a.userRepo.GetByEmail(ctx, email)

	if err != nil {
		a.uow.RollBack()
		return uuid.Nil, err
	}

	a.uow.Commit()
	return user.Id, nil
}
