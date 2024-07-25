package database

import (
	// "catalog-bot-api/structs"

	"catalog-bot-api/structs"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		Shop{},
		CatalogItem{},
	)
}

// CRUD operations

func CreateShop(db *gorm.DB, data Shop) error {
	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateItem(db *gorm.DB, data CatalogItem) error {
	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllItems(db *gorm.DB, shopID int) []structs.CatalogItemResponse {
	var catalogItems []structs.CatalogItemResponse
	db.Model(&CatalogItem{}).Where("shop_id = ?", shopID).Find(&structs.CatalogItemResponse{}).Find(&catalogItems)

	return catalogItems
}

func GetShopTitle(db *gorm.DB, shopID int) string {
	var shopTitle string
	db.Raw("SELECT title FROM shops WHERE id = ?", shopID).Scan(&shopTitle)
	return shopTitle
}

func GetOneItem(db *gorm.DB, itemID int) structs.CatalogItemResponse {
	var catalogItem structs.CatalogItemResponse
	db.Model(&CatalogItem{}).Where("id = ?", itemID).Find(&structs.CatalogItemResponse{}).Find(&catalogItem)

	return catalogItem
}