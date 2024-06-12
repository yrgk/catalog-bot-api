package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"catalog-bot-api/database"
)

func RunServer(db *gorm.DB) {
	app := fiber.New()

	// catalog := app.Group("/catalog")

	// Fetching all catalogs in one shop
	app.Get(":shopID/catalogs", func(c *fiber.Ctx) error {
		shopID, err := strconv.Atoi(c.Params("shopID"))
		if err != nil {
			return c.JSON(BadResponse{
				false,
				"incorrect shop id",
			})
		}
		catalogs, err := database.GetAllCatalogs(db, shopID)
		if err != nil {
			return c.JSON(BadResponse{
				false,
				"not found",
			})
		}

		return c.JSON(catalogs)
	})

	// Fetching one item from catalog
	// catalog.Get("/:id", func(c *fiber.Ctx) error {
	// 	id, err := strconv.Atoi(c.Params("id"))
	// 	if err != nil {
	// 		return c.JSON(BadResponse{
	// 			false,
	// 			"incorect id",
	// 		})
	// 	}
	// 	item := database.GetOneItem(db, id)

	// 	return c.JSON(item)
	// })

	// catalog.Get("/:id", func(c *fiber.Ctx) error {
	// 	id, err := strconv.Atoi(c.Params("id"))
	// 	if err != nil {
	// 		return c.JSON(BadResponse{
	// 			false,
	// 			"incorect id",
	// 		})
	// 	}
	// 	item := database.GetOneItem(db, id)

	// 	return c.JSON(item)
	// })

	// catalog.Post("/", func(c *fiber.Ctx) error {

	// })

	app.Listen(":8080")
}
