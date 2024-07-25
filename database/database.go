package database

import (
	"catalog-bot-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func InitDb(config config.Config) *gorm.DB {
	dsn := config.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}