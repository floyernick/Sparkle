package posts

import (
	"Sparkle/gateways/common/database"
)

type PostsGateway struct {
	db database.DB
}

func New(db database.DB) PostsGateway {
	return PostsGateway{db}
}
