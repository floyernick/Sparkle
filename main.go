package main

import (
	"flag"

	"Sparkle/config"
	"Sparkle/database"
	"Sparkle/server"
	"Sparkle/tools/logger"
)

func main() {

	flag.Parse()

	cfg, err := config.LoadConfig()

	if err != nil {
		logger.Error(err)
	}

	db, err := database.Sparkle(cfg.Database)

	if err != nil {
		logger.Error(err)
	}

	err = server.Sparkle(cfg.Server, db)

	if err != nil {
		logger.Error(err)
	}

}
