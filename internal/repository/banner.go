package repository

import (
	"catalog-bot-api/internal/models"
	"catalog-bot-api/pkg/postgres"
)

func GetBanners(shopId int) []models.Banner {
	var banners []models.Banner
	postgres.DB.Where("shop_id = ?", shopId).Find(&banners)

	return banners
}

func GetBannerById(id int) models.Banner {
	var banner models.Banner
	postgres.DB.First(&banner)

	return banner
}