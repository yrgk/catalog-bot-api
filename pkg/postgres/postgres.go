package postgres

import (
	"catalog-bot-api/internal/config"
	"catalog-bot-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := config.Config.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		models.Shop{},
		models.CatalogItem{},
		models.Banner{},
	)

	DB = db
}
