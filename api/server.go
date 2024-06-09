package api

import "github.com/gofiber/fiber/v2"


func RunServer() {
	app := fiber.New()


	app.Get("/", func(c *fiber.Ctx) error {
		
	})
}