package main

import (
	"Sparkle/config"
	"Sparkle/delivery"
	"Sparkle/gateways/common/cache"
	"Sparkle/gateways/common/database"
	"Sparkle/tools/logger"

	likesGateway "Sparkle/gateways/likes"
	locationsGateway "Sparkle/gateways/locations"
	postsGateway "Sparkle/gateways/posts"
	usersGateway "Sparkle/gateways/users"

	likesController "Sparkle/delivery/likes"
	locationsController "Sparkle/delivery/locations"
	postsController "Sparkle/delivery/posts"
	usersController "Sparkle/delivery/users"

	likesService "Sparkle/services/likes"
	locationsService "Sparkle/services/locations"
	postsService "Sparkle/services/posts"
	usersService "Sparkle/services/users"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		logger.Fatal(err)
	}

	err = logger.Setup(cfg.Logger)

	if err != nil {
		logger.Fatal(err)
	}

	db, err := database.Init(cfg.Database)

	if err != nil {
		logger.Fatal(err)
	}

	cache, err := cache.Init(cfg.Cache)

	if err != nil {
		logger.Fatal(err)
	}

	usersGateway := usersGateway.New(db)
	postsGateway := postsGateway.New(db)
	likesGateway := likesGateway.New(db)
	locationsGateway := locationsGateway.New(db, cache)

	usersService := usersService.New(usersGateway, postsGateway)
	postsService := postsService.New(usersGateway, postsGateway)
	likesService := likesService.New(usersGateway, postsGateway, likesGateway)
	locationsService := locationsService.New(usersGateway, locationsGateway)

	usersController := usersController.New(usersService)
	postsController := postsController.New(postsService)
	likesController := likesController.New(likesService)
	locationsController := locationsController.New(locationsService)

	err = delivery.Serve(cfg.Server, usersController, postsController, likesController, locationsController)

	if err != nil {
		logger.Fatal(err)
	}

}
