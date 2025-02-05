package interfaces

import (
	"context"
	"net/http"
)

type AuthAPI interface {
	LogInHandler(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(mux *http.ServeMux)
}

type AuthHandler interface {
	LogIn(ctx context.Context, uid string) (string, error)
}
