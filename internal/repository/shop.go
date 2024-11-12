package repository

import (
	"catalog-bot-api/internal/models"
	"catalog-bot-api/pkg/postgres"
	"fmt"
)

// TODO: change to native sql
func CreateShop(shopData models.Shop) error {
	result := postgres.DB.Create(&shopData)
	if result.Error != nil {
		return result.Error
	}
	return nil
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
