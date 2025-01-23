package interfaces

type AuthHandler interface {
	SignUp(email string, password string) (string, error)
	SignIn(email string, password string) (string, error)
}
