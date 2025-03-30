package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: os.Getenv("DATABASE_URL"),
	}
}

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found.")
		return
	}
	log.Println(".env file loaded.")
}
