package interfaces

import (
	"context"

	"github.com/Quan0308/main-api/models"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}
