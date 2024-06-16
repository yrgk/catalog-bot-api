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

func GetAllItems(db *gorm.DB, catalogID int) ([]structs.CatalogItemReponse, error) {
	var catalogItems []structs.CatalogItemReponse
	result := db.Model(&CatalogItem{}).Where("catalog_id = ?", catalogID).Find(&structs.CatalogItemReponse{}).Find(&catalogItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return catalogItems, nil
}

func GetAllCatalogs(db *gorm.DB, shopID int) ([]Catalog, error) {
	var catalogs []Catalog
	// err := db.Where("shop_id = ?", shopID).Find(&catalogs)
	db.Where("shop_id = ?", shopID).Find(&catalogs)
	// if err != nil {
	// 	return nil, err.Error
	// }

	return catalogs, nil
}

func GetOneItem(db *gorm.DB, id int) CatalogItem {
	var catalogItem CatalogItem

	db.Find(&catalogItem, id)

	return catalogItem
}