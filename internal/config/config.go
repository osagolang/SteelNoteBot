package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN not found in .env")
	}

	return token
}
