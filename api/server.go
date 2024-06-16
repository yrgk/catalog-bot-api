package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"catalog-bot-api/database"
	"catalog-bot-api/structs"
)

func RunServer(db *gorm.DB) {
	app := fiber.New()

	shop := app.Group("/shop")
	catalog := app.Group("/catalog")
	item := app.Group("/item")

	// Fetching info about the shop
	// shop.Get("/:shopID/info", func(c *fiber.Ctx) error {
	// 	shopID, err := strconv.Atoi(c.Params("shopID"))
	// 	if err != nil {
	// 		return c.JSON(structs.SimpleResponse{
	// 			IsOk: false,
	// 			Message: "invalid shop id",
	// 			StatusCode: 404,
	// 		})
	// 	}
	// 	catalogs := database.GetAllCatalogs(db, shopID)

	// 	return c.JSON(catalogs)
	// })

	// Fetching all catalogs in one shop
	shop.Get("/:shopID/catalog/all", func(c *fiber.Ctx) error {
		shopID, err := strconv.Atoi(c.Params("shopID"))
		if err != nil {
			return c.JSON(structs.SimpleResponse{
				IsOk: false,
				Message: "invalid shop id",
				StatusCode: 404,
			})
		}
		catalogs := database.GetAllCatalogs(db, shopID)

		return c.JSON(catalogs)
	})

	// Fetching all items in one shop
	catalog.Get("/:catalogID", func(c *fiber.Ctx) error {
		catalogID, err := strconv.Atoi(c.Params("catalogID"))
		if err != nil {
			return c.JSON(structs.SimpleResponse{
				IsOk: false,
				Message: "invalid catalog id",
				StatusCode: 404,
			})
		}
		catalogs := database.GetAllItems(db, catalogID)

		return c.JSON(catalogs)
	})

	// Fetching one item from catalog
	item.Get("/:itemID", func(c *fiber.Ctx) error {
		itemID, err := strconv.Atoi(c.Params("itemID"))
		if err != nil {
			return c.JSON(structs.SimpleResponse{
				IsOk: false,
				Message: "invalid item id",
				StatusCode: 404,
			})
		}
		item := database.GetOneItem(db, itemID)

		return c.JSON(item)
	})

	// catalog.Post("/", func(c *fiber.Ctx) error {

	// })

	app.Listen(":8080")
}
