package database

import (
	"catalog-bot-api/structs"
	"fmt"

	"gorm.io/gorm"
)

func GetOneItem(db *gorm.DB, itemID int) structs.OneCatalogItemResponse {
	var catalogItem structs.OneCatalogItemResponse
	query := fmt.Sprintf("SELECT catalog_items.id, catalog_items.title, catalog_items.price, catalog_items.description, catalog_items.cover_url, shops.currency FROM catalog_items LEFT JOIN shops ON shops.id = catalog_items.shop_id WHERE catalog_items.id = %d", itemID)
	db.Raw(query).Find(&catalogItem)

	return catalogItem
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
