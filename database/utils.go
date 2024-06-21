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

// func GetShopInfo(db *gorm.DB, shopID int)

func GetAllItems(db *gorm.DB, shopID int) []structs.CatalogItemResponse {
	var catalogItems []structs.CatalogItemResponse
	db.Model(&CatalogItem{}).Where("shop_id = ?", shopID).Find(&structs.CatalogItemResponse{}).Find(&catalogItems)

	return catalogItems
}

func GetOneItem(db *gorm.DB, itemID int) structs.OneCatalogItemResponse {
	var catalogItem structs.OneCatalogItemResponse
	db.Model(&CatalogItem{}).Where("id = ?", itemID).Find(&structs.OneCatalogItemResponse{}).Find(&catalogItem)

	return catalogItem
}