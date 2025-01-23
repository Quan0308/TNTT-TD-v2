package api

import (
	"net/http"

	"github.com/Quan0308/main-api/handlers"
)

func RegisterRoutes(server *http.ServeMux) {
	authHandler := handlers.NewAuthHandler()
	authAPI := NewAuthAPI(authHandler)

	authAPI.RegisterRoutes(server)
}
