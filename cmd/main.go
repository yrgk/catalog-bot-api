package main

import (
	"catalog-bot-api/internal/config"
	"catalog-bot-api/internal/handlers"
	"catalog-bot-api/pkg/postgres"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.GetConfig()
	postgres.ConnectDb()

	app := fiber.New()

	handlers.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%s", config.Config.Port)))
}
