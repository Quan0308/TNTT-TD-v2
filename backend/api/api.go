package api

import (
	"github.com/Quan0308/main-api/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(server *mux.Router) {
	authHandler := handlers.NewAuthHandler()
	authAPI := NewAuthAPI(authHandler)

	authAPI.RegisterRoutes(server)
}
