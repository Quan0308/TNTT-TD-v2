package interfaces

import (
	"context"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/Quan0308/main-api/models"
)

type AuthAPI interface {
	LogInHandler(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(mux *http.ServeMux)
}

type AuthHandler interface {
	LogIn(ctx context.Context, uid string) (string, error)
}

type FirebaseAuthService interface {
	GenerateCustomToken(ctx context.Context, uid string, claims *models.CurrentUser) (string, error)
	VerifyIdToken(ctx context.Context, idToken string) (*auth.Token, error)
}
