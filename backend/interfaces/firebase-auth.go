package interfaces

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/Quan0308/main-api/models"
)

type FirebaseAuthService interface {
	GenerateCustomToken(ctx context.Context, uid string, claims *models.CurrentUser) (string, error)
	VerifyIdToken(ctx context.Context, idToken string) (*auth.Token, error)
}
