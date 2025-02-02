package interfaces

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type AuthAPI interface {
	LogInHandler(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(mux *http.ServeMux)
}

type AuthHandler interface {
	LogIn(ctx context.Context, email string, password string) (uuid.UUID, error)
}
