package database

import (
	"time"
)

type (
	Shop struct {
		ID             uint      `gorm:"primarykey"`
		CreatedAt      time.Time `json:"created_at"`
		Title          string    `json:"title"`
		TelegramUserID int       `json:"telegram_user_id"`
	}

	CatalogItem struct {
		ID          uint      `gorm:"primarykey"`
		CreatedAt   time.Time `json:"created_at"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Price       int       `json:"price"`
		CoverUrl    string    `json:"cover_url"`
		Currency    string    `json:"currency"`
		ShopID      int       `json:"shop_id"`
	}
)
