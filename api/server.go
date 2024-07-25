package api

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"

	"catalog-bot-api/database"
	"catalog-bot-api/structs"
)

func RunServer(db *gorm.DB) {
	var PORT string
	strPort := os.Getenv("PORT")
	if strPort == "" {
		PORT = "8080"
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	shop := app.Group("/shop")
	catalog := app.Group("/catalog")
	item := app.Group("/item")

	// Creating the shop
	shop.Post("/new", func(c *fiber.Ctx) error {
		payload := new(structs.ShopCreateRequest)

		if err := c.BodyParser(payload); err != nil {
			return err
		}

		// TODO: Add payment checking

		newShop := database.Shop{
			CreatedAt: time.Now(),
			Title: payload.Title,
			TelegramUserID: payload.TelegramUserID,
		}

		if err := database.CreateShop(db, newShop); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(newShop)
	})

	// Fetching all items in one shop
	catalog.Get("/:shopID", func(c *fiber.Ctx) error {
		shopID, err := c.ParamsInt("shopID")
		log.Println(shopID)
		if err != nil {
			var result []structs.CatalogResponse
			return c.Status(fiber.StatusNotFound).JSON(result)
		}
		catalogs := database.GetAllItems(db, shopID)
		if len(catalogs) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(catalogs)
		}
		shopTitle := database.GetShopTitle(db, shopID)
		result := structs.CatalogResponse{
			ShopTitle: shopTitle,
			Items: catalogs,
		}

		return c.Status(fiber.StatusOK).JSON(result)
	})

	// Fetching one item from catalog
	item.Get("/:itemID", func(c *fiber.Ctx) error {
		itemID, err := c.ParamsInt("itemID")
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		item := database.GetOneItem(db, itemID)

		return c.JSON(item)
	})

	item.Post("/new", func(c *fiber.Ctx) error {
		payload := new(structs.CatalogItemRequest)

		if err := c.BodyParser(payload); err != nil {
			return err
		}

		newItem := database.CatalogItem{
			CreatedAt: time.Now(),
			Title: payload.Title,
			Description: payload.Description,
			Price: payload.Price,
			CoverUrl: payload.CoverUrl,
			Currency: payload.Currency,
			ShopID: payload.ShopID,
		}

		if err := database.CreateItem(db, newItem); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Listen(fmt.Sprintf(":%s", PORT))
}
