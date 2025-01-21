package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string

	DB_USER     string
	DB_PASSWORD string
	DB_ADDRESS  string
	DB_NAME     string
}

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	return Config{
		PORT:        getEnv("PORT", "8080"),
		DB_USER:     getEnv("DB_USER", "root"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "1"),
		DB_NAME:     getEnv("DB_NAME", "system_manage"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var Envs = initConfig()
