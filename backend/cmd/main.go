package main

import (
	"log"
	"net/http"

	"github.com/Quan0308/main-api/api"
	"github.com/Quan0308/main-api/config"
	"github.com/Quan0308/main-api/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	port := config.Envs.PORT
	addr := ":" + port
	server := mux.NewRouter()

	server.Use(middlewares.LoggingMiddleware)
	api.RegisterRoutes(server)

	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(addr, server))
}
