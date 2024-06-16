package database

import (
	"gorm.io/gorm"
)

type (
	Shop struct {
		gorm.Model
		Title          string `json:"title"`
		TelegramUserID int    `json:"telegram_user_id"`
	}

	Catalog struct {
		gorm.Model
		Title    string `json:"title"`
		CoverUrl string `json:"cover_url"`
		ShopID   int    `json:"shop_id"`
	}

	CatalogItem struct {
		gorm.Model
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       int    `json:"price"`
		CoverUrl    string `json:"cover_url"`
		Currency    string `json:"currency"`
		CatalogID   int    `json:"catalog_id"`
	}
)
