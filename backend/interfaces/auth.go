package interfaces

import (
	"context"

	"github.com/google/uuid"
)

type AuthHandler interface {
	LogIn(ctx context.Context, email string, password string) (uuid.UUID, error)
}
