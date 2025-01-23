package handlers

import "github.com/Quan0308/main-api/core/interfaces"

type authHandlerImpl struct{}

func NewAuthHandler() interfaces.AuthHandler {
	return &authHandlerImpl{}
}

func (a *authHandlerImpl) SignUp(email string, password string) (string, error) {
	return "success", nil
}

func (a *authHandlerImpl) SignIn(email string, password string) (string, error) {
	return "success", nil
}
