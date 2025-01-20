package handlers

type authHandlerImpl struct{}

func NewAuthHandler() AuthHandler {
	return &authHandlerImpl{}
}

func (a *authHandlerImpl) SignUp(email string, password string) (string, error) {
	return "success", nil
}

func (a *authHandlerImpl) SignIn(email string, password string) (string, error) {
	return "success", nil
}
