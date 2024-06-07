package database

import (
	"gorm.io/gorm"
)

type Catalog struct {
	gorm.Model
	Title          string
	UserTelegramID int
}

type CatalogItem struct {
	gorm.Model
	Title       string
	Description string
	Price       int
	Currency    string
	CatalogID   int
	CategoryID  int
}
