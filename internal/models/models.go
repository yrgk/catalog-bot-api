package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Shop struct {
		ID             uint   `json:"id" gorm:"primarykey"`
		Title          string `json:"title" gorm:"not null"`
		ExpirationDate string `json:"expiration_date"`
		Currency       string `json:"currency"`
		ChannelUrl     string `json:"channel_url"`
		TelegramUserID int    `json:"telegram_user_id" gorm:"not null"`
	}

	CatalogItem struct {
		ID            uint      `json:"id" gorm:"primarykey"`
		CreatedAt     time.Time `json:"created_at" gorm:"not null"`
		Title         string    `json:"title" gorm:"not null"`
		Description   string    `json:"description"`
		Price         float32   `json:"price" gorm:"not null"`
		DiscountPrice float32   `json:"discount_price"`
		CoverUrl      string    `json:"cover_url"`
		ShopID        uint      `json:"shop_id" gorm:"not null"`
	}

	Banner struct {
		ID          uint   `json:"id" gorm:"primarykey"`
		Description string `json:"description"`
		CoverUrl    string `json:"cover_url"`
		ShopID      uint   `json:"shop_id" gorm:"not null"`
	}

	Photo struct {
		ID       uint   `gorm:"primarykey"`
		CoverUrl string `json:"cover_url"`
	}

	Order struct {
		gorm.Model
		UserId int
		Units  []Unit `gorm:"foreignKey:OrderId"`
	}

	Unit struct {
		gorm.Model
		Title   string
		OrderId int
	}
)
