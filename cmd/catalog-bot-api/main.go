package main

import (
	"catalog-bot-api/api"
	"catalog-bot-api/config"
	"catalog-bot-api/database"
)

func main() {
	config := config.GetConfig()
	db := database.InitDb(config)
	database.Migrate(db)
	api.RunServer(db)
}