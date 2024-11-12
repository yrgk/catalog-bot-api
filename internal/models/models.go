package models

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
		ID          uint      `json:"id" gorm:"primarykey"`
		CreatedAt   time.Time `json:"created_at" gorm:"not null"`
		Title       string    `json:"title" gorm:"not null"`
		Description string    `json:"description"`
		Price       float32   `json:"price" gorm:"not null"`
		CoverUrl    string    `json:"cover_url"`
		Discount    uint      `json:"discount"`
		ShopID      uint      `json:"shop_id" gorm:"not null"`
	}

	Banner struct {
		ID          uint      `json:"id" gorm:"primarykey"`
		Description string    `json:"description"`
		CoverUrl    string    `json:"cover_url"`
		ShopID      uint      `json:"shop_id" gorm:"not null"`
	}

	Photo struct {
		ID       uint   `gorm:"primarykey"`
		CoverUrl string `json:"cover_url"`
	}
)
