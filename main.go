package main

import (
	"Sparkle/config"
	"Sparkle/database"
	"Sparkle/server"
	"Sparkle/tools/logger"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		logger.Error(err)
	}

	db, err := database.Init(cfg.Database)

	if err != nil {
		logger.Error(err)
	}

	err = server.Init(cfg.Server, db)

	if err != nil {
		logger.Error(err)
	}

}
