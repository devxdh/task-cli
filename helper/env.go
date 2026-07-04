package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}

	DB_ENV := "DATABASE_URL"

	DB_URL := os.Getenv(DB_ENV)
	if len(DB_URL) == 0 {
		log.Fatal("Error loading %s from .env: %v", DB_ENV, err)
	}

	return map[string]string{
		DB_ENV: DB_URL,
	}
}
