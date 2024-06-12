package database

import (
	"gorm.io/gorm"
)

type (
	Shop struct {
		gorm.Model
		Title          string
		TelegramUserID int
	}
	Catalog struct {
		gorm.Model
		Title    string
		CoverUrl string
		ShopID   int
	}

	CatalogItem struct {
		gorm.Model
		Title       string
		Description string
		Price       int
		CoverUrl    string
		Currency    string
		CatalogID   int
		CategoryID  int
	}
)
