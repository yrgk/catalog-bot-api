package main

import (
	"catalog-bot-api/internal/config"
	"catalog-bot-api/internal/handlers"
	"catalog-bot-api/pkg/postgres"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.GetConfig()
	postgres.ConnectDb()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	handlers.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s",config.Config.Port)))
	// err := app.ListenTLS(":443", "/etc/letsencrypt/live/catalogio.space/fullchain.pem", "/etc/letsencrypt/live/catalogio.space/privkey.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
