package main

import (
	"catalog-bot-api/api"
	"catalog-bot-api/database"
)

func main() {
	db := database.GetDb()
	// database.Migrate(db)
	api.RunServer(db)
}