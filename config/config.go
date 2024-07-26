package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN  string
	Port string
}

func GetConfig() Config {
	if err := godotenv.Load("./../../"); err != nil {
		log.Fatalf(".env file not found: %s", err)
	}

	// Getting values from environment file
	DSN := os.Getenv("DSN")
	port := os.Getenv("PORT")

	return Config{
		DSN:  DSN,
		Port: port,
	}
}
