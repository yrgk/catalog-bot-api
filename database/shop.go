package database

import (
	"catalog-bot-api/structs"
	"fmt"

	"gorm.io/gorm"
)

// TODO: change to native sql
func CreateShop(db *gorm.DB, data Shop) error {
	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetMyShops(db *gorm.DB, userId string) []structs.MyShopsResponse {
	var shops []structs.MyShopsResponse
	query := fmt.Sprintf("SELECT shops.id, shops.title FROM shops WHERE telegram_user_id = %s", userId)
	db.Raw(query).Find(&shops)

	return shops
}

func GetShopData(db *gorm.DB, shopID int) structs.ShopData {
	var shopData structs.ShopData
	db.Raw("SELECT title, currency FROM shops WHERE id = ?", shopID).Scan(&shopData)
	return shopData
}
