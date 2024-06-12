package database

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		Shop{},
		Catalog{},
		CatalogItem{},
	)
}

// func GetAllItems(db *gorm.DB) [] {
// 	var catalogItems []CatalogItem

// 	db.Where()
// }
func GetAllCatalogs(db *gorm.DB, shopID int) ([]Catalog, error) {
	var catalogs []Catalog
	err := db.Where("shop_id = ?", shopID).Find(&catalogs)
	if err != nil {
		return nil, err.Error
	}

	return catalogs, nil
}

func GetOneItem(db *gorm.DB, id int) CatalogItem {
	var catalogItem CatalogItem

	db.Find(&catalogItem, id)

	return catalogItem
}