package interfaces

import (
	"github.com/google/uuid"
)

type AuthHandler interface {
	SignUp(email string, password string) (string, error)
	SignIn(email string, password string) (uuid.UUID, error)
}
