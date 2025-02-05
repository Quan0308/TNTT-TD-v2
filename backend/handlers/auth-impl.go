package handlers

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/Quan0308/main-api/interfaces"
)

type authHandlerImpl struct {
	uow                interfaces.UnitOfWork
	userRepo           interfaces.UserRepository
	firebaseAuthServie interfaces.FirebaseAuthService
}

func NewAuthHandler(uow interfaces.UnitOfWork, userRepo interfaces.UserRepository, firebaseAuthServie interfaces.FirebaseAuthService) interfaces.AuthHandler {
	return &authHandlerImpl{
		uow:                uow,
		userRepo:           userRepo,
		firebaseAuthServie: firebaseAuthServie,
	}
}

func (a *authHandlerImpl) LogIn(ctx context.Context, uid string) (string, error) {
	err := a.uow.Begin()
	if err != nil {
		return "", err
	}

	currentUser, err := a.userRepo.GetCurrentUserByUid(ctx, uid)

	if err != nil {
		a.uow.RollBack()
		return "", err
	}

	accessToken, err := a.firebaseAuthServie.GenerateCustomToken(ctx, uid, currentUser)
	if err != nil {
		a.uow.RollBack()
		return "", err
	}

	a.uow.Commit()
	return accessToken, nil
}

func (a *authHandlerImpl) VerifyIdToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return a.firebaseAuthServie.VerifyIdToken(ctx, idToken)
}
