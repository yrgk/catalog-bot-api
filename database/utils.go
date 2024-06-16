package database

import (
	// "catalog-bot-api/structs"

	"catalog-bot-api/structs"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		Shop{},
		Catalog{},
		CatalogItem{},
	)
}

// func GetShopInfo(db *gorm.DB, shopID int)

func GetAllItems(db *gorm.DB, catalogID int) []structs.CatalogItemResponse {
	var catalogItems []structs.CatalogItemResponse
	db.Model(&CatalogItem{}).Where("catalog_id = ?", catalogID).Find(&structs.CatalogItemResponse{}).Find(&catalogItems)

	return catalogItems
}

func GetAllCatalogs(db *gorm.DB, shopID int) []structs.CatalogResponse {
// func GetAllCatalogs(db *gorm.DB, shopID int) []Catalog {
	// var catalogs []Catalog
	var catalogs []structs.CatalogResponse
	db.Model(&Catalog{}).Where("shop_id = ?", shopID).Find(&structs.CatalogResponse{}).Find(&catalogs)

	return catalogs
}

func GetOneItem(db *gorm.DB, itemID int) structs.OneCatalogItemResponse {
	var catalogItem structs.OneCatalogItemResponse
	db.Model(&CatalogItem{}).Where("id = ?", itemID).Find(&structs.OneCatalogItemResponse{}).Find(&catalogItem)

	return catalogItem
}