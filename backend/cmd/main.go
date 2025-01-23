package main

import (
	"log"
	"net/http"

	"github.com/Quan0308/main-api/api"
	"github.com/Quan0308/main-api/config"
	"github.com/Quan0308/main-api/core/middlewares"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := config.Envs.CONNECTION_STRING
	db, err := sqlx.Connect(config.Envs.DRIVER, dsn)
	if err != nil {
		log.Fatal("Fail to connect to database: ", err)
	}
	defer db.Close()

	v2Router := http.NewServeMux()
	api.RegisterRoutes(v2Router, db)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/api/v2/", http.StripPrefix("/api/v2", v2Router))

	server := http.Server{
		Addr:    config.Envs.ADDR,
		Handler: middlewares.Logging(mainRouter),
	}

	log.Println("Listening on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
