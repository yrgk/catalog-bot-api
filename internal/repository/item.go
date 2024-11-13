package repository

import (
	"catalog-bot-api/internal/models"
	"catalog-bot-api/pkg/postgres"
	"fmt"
)

func GetOneItem(itemID int) models.ItemResponse {
	var catalogItem models.ItemResponse
	query := fmt.Sprintf("SELECT catalog_items.id, catalog_items.title, catalog_items.price, catalog_items.discount, catalog_items.description, catalog_items.cover_url, shops.currency FROM catalog_items LEFT JOIN shops ON shops.id = catalog_items.shop_id WHERE catalog_items.id = %d", itemID)
	postgres.DB.Raw(query).Find(&catalogItem)

	return catalogItem
}

func CreateItem(data models.CatalogItem) error {
	result := postgres.DB.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetItemsInOneShop(shopID int) []models.CatalogItemResponse {
	var catalogItems []models.CatalogItemResponse
	postgres.DB.Model(&models.CatalogItem{}).Where("shop_id = ?", shopID).Find(&models.CatalogItemResponse{}).Find(&catalogItems)

	return catalogItems
}
