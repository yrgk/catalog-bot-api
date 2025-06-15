package handlers

import (
	"catalog-bot-api/internal/models"
	"catalog-bot-api/internal/repository"
	"catalog-bot-api/pkg/postgres"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateShopHandler(c *fiber.Ctx) error {
	payload := new(models.ShopCreateRequest)

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	dateNextMonth := time.Now().AddDate(0, 1, 0).Format("2006-01-02 15:04:05")

	newShop := models.Shop{
		Title:          payload.Title,
		ExpirationDate: dateNextMonth,
		Currency:       payload.Currency,
		// ChannelUrl:     payload.ChannelUrl,
		TelegramUserID: payload.TelegramUserID,
	}

	response, err := repository.CreateShop(newShop)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(response)
}

func GetMyShopsHandler(c *fiber.Ctx) error {
	userId := c.Query("user_id")

	shops := repository.GetMyShops(userId)
	if len(shops) == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(shops)
}

func GetCatalogHandler(c *fiber.Ctx) error {
	shopId, err := c.ParamsInt("shopId")
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	catalogs := repository.GetItemsInOneShop(shopId)
	if len(catalogs) == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	shopData := repository.GetShopData(shopId)

	banners := repository.GetBanners(shopId)

	result := models.CatalogResponse{
		ShopTitle: shopData.Title,
		Currency:  shopData.Currency,
		Banners:   banners,
		Items:     catalogs,
	}

	return c.JSON(result)
}

func GetItemHandler(c *fiber.Ctx) error {
	itemID, err := c.ParamsInt("itemID")
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	item := repository.GetOneItem(itemID)

	return c.JSON(item)
}

func CreateItemHandler(c *fiber.Ctx) error {
	body := new(models.CatalogItemRequest)

	if err := c.BodyParser(body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newItem := models.CatalogItem{
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
	}

	if err := repository.CreateItem(newItem); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetBannerHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("bannerId")
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	banner := repository.GetBannerById(id)
	if banner.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(banner)
}

func CreateOrderHandler(c *fiber.Ctx) error {
	var body models.CreateOrderRequest
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := repository.CreateOrder(body); err != nil {
		return c.SendStatus(fiber.StatusConflict)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetOrderHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	response := repository.GetOrder(id)

	return c.JSON(response)
}

func GetOrdersByUser(c *fiber.Ctx) error {
    userIdStr := c.Query("userId")
    userId, err := strconv.Atoi(userIdStr)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid userId"})
    }

    var orders []models.Order
    if err := postgres.DB.Preload("Units").Where("user_id = ?", userId).Find(&orders).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch orders"})
    }

    var response []models.OrderListResponse
    for _, order := range orders {
        units := make([]models.UnitResponse, len(order.Units))
        for i, u := range order.Units {
            units[i] = models.UnitResponse{
                ID:    u.ID,
                Title: u.Title,
				Price: u.Price,
            }
        }

        response = append(response, models.OrderListResponse{
            OrderID: order.ID,
			State: order.State,
            Units:   units,
        })
    }

    return c.JSON(response)
}