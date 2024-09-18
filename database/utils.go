package database

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		Shop{},
		CatalogItem{},
	)
}
