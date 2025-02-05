package interfaces

import (
	"context"

	"github.com/Quan0308/main-api/models"
)

type UserRepository interface {
	GetCurrentUserByUid(ctx context.Context, email string) (*models.CurrentUser, error)
}
