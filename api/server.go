package api

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"catalog-bot-api/database"
	"catalog-bot-api/structs"
)

func RunServer(db *gorm.DB) {
	app := fiber.New()

	// shop := app.Group("/shop")
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

	// Fetching all items in one shop
	catalog.Get("/:shopID", func(c *fiber.Ctx) error {
		shopID, err := c.ParamsInt("shopID")
		log.Println(shopID)
		if err != nil {
			var result []structs.CatalogItemResponse
			return c.Status(fiber.StatusNotFound).JSON(result)
		}
		catalogs := database.GetAllItems(db, shopID)
		if len(catalogs) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(catalogs)
		}
		return c.Status(fiber.StatusOK).JSON(catalogs)


	})

	// Fetching one item from catalog
	item.Get("/:itemID", func(c *fiber.Ctx) error {
		itemID, err := c.ParamsInt("itemID")
		if err != nil {
			return c.JSON(structs.SimpleResponse{
				IsOk: false,
				Message: "invalid item id",
				StatusCode: 404,
			})
		}
		item := database.GetOneItem(db, itemID)

		return c.Status(fiber.StatusOK).JSON(item)
	})

	item.Post("/", func(c *fiber.Ctx) error {
		payload := new(structs.CatalogItemRequest)

		if err := c.BodyParser(payload); err != nil {
			return err
		}

		log.Println(payload.Title)

		newItem := database.CatalogItem{
			CreatedAt: time.Now(),
			Title: payload.Title,
			Description: payload.Description,
			Price: payload.Price,
			CoverUrl: payload.CoverUrl,
			Currency: payload.Currency,
			ShopID: 1,
		}

		return c.Status(fiber.StatusOK).JSON(newItem)
	})

	app.Listen(":8080")
}
