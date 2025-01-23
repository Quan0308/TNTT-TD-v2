package api

import (
	"net/http"

	"github.com/Quan0308/main-api/handlers"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(server *http.ServeMux, db *sqlx.DB) {
	authHandler := handlers.NewAuthHandler(db)
	authAPI := NewAuthAPI(authHandler)

	authAPI.RegisterRoutes(server)
}
