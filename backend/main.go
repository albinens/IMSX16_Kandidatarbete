package main

import (
	"example.com/m/v2/api"
	"example.com/m/v2/database"
	"example.com/m/v2/env"
	"example.com/m/v2/logger"
	"example.com/m/v2/seeder"
	"example.com/m/v2/utils"
)

func main() {
	env.Load()
	logger.Configure()
	if err := database.InitSQL(); err != nil {
		utils.LogFatal("Failed to initialize SQL: ", err)
	}
	if err := seeder.SeedDevelopmentData(); err != nil {
		utils.LogFatal("Failed to seed development data: ", err)
	}

	api.Init()
}
