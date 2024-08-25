package database

import (
	"time"
)

type (
	Shop struct {
		ID             uint      `json:"id" gorm:"primarykey"`
		CreatedAt      time.Time `json:"created_at" gorm:"not null"`
		Title          string    `json:"title" gorm:"not null; unique"`
		ExpirationDate time.Time `json:"expiration_date"`
		Currency       string    `json:"currency"`
		ChannelUrl     string    `json:"channel_url"`
		TelegramUserID int       `json:"telegram_user_id" gorm:"not null"`
	}

	CatalogItem struct {
		ID          uint      `gorm:"primarykey"`
		CreatedAt   time.Time `json:"created_at" gorm:"not null"`
		Title       string    `json:"title" gorm:"not null"`
		Description string    `json:"description"`
		Price       float32   `json:"price" gorm:"not null"`
		CoverUrl    string    `json:"cover_url"`
		// Currency    string    `json:"currency" gorm:"not null"`
		ShopID int `json:"shop_id" gorm:"not null"`
	}
)
