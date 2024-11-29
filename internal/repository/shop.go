package repository

import (
	"catalog-bot-api/internal/models"
	"catalog-bot-api/pkg/postgres"
	"fmt"
)

func CreateShop(shopData models.Shop) (models.CreateShopResponse, error) {
	var response models.CreateShopResponse
	query := fmt.Sprintf("INSERT INTO shops (title, currency, telegram_user_id, expiration_date) VALUES ('%s', '%s', '%d', '%s') RETURNING id, expiration_date", shopData.Title, shopData.Currency, shopData.TelegramUserID, shopData.ExpirationDate)
	postgres.DB.Raw(query).Scan(&response)
	// if result.Error != nil {
	// 	return models.CreateShopResponse{}, result.Error
	// }
	return response, nil
}

func GetMyShops(userId string) []models.MyShopsResponse {
	var shops []models.MyShopsResponse
	query := fmt.Sprintf("SELECT shops.id, shops.title FROM shops WHERE telegram_user_id = %s", userId)
	postgres.DB.Raw(query).Find(&shops)

	return shops
}

func GetShopData(shopID int) models.ShopData {
	var shopData models.ShopData
	postgres.DB.Raw("SELECT title, currency FROM shops WHERE id = ?", shopID).Scan(&shopData)
	return shopData
}
