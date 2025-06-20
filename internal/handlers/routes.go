package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Setting up CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Setting up logger
	app.Use(logger.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/catalog/new", CreateShopHandler)
	v1.Post("/item/new", CreateItemHandler)

	v1.Get("/catalog/my", GetMyShopsHandler)
	v1.Get("/catalog/:shopId", GetCatalogHandler)
	v1.Get("/item/:itemID", GetItemHandler)
	v1.Get("/banner/:bannerId", GetBannerHandler)

	v1.Post("/order", CreateOrderHandler)
	v1.Get("/order/list", GetOrdersByUser)
	v1.Get("/order/:id", GetOrderHandler)
}
