package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ADDR               string
	DRIVER             string
	CONNECTION_STRING  string
	SI_SERVICE_ACCOUNT string
}

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	return Config{
		ADDR:               ":" + getEnv("PORT", "8080"),
		DRIVER:             getEnv("DRIVER", ""),
		CONNECTION_STRING:  getEnv("CONNECTION_STRING", ""),
		SI_SERVICE_ACCOUNT: getEnv("SI_SERVICE_ACCOUNT", ""),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var Envs = initConfig()
