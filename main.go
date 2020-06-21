package main

import (
	"flag"

	"Init/config"
	"Init/database"
	"Init/server"
	"Init/tools/logger"
)

func main() {

	flag.Parse()

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
