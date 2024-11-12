package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	DSN  string
	Port string
}

var Config ConfigStruct

func GetConfig() {
	// if err := godotenv.Load("../../.env"); err != nil {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found: %s", err)
	}

	// Getting values from environment file
	DSN := os.Getenv("DSN")
	port := os.Getenv("PORT")

	Config = ConfigStruct{
		DSN:  DSN,
		Port: port,
	}
}
