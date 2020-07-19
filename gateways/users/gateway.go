package users

import (
	"Sparkle/gateways/common/database"
)

type UsersGateway struct {
	db database.DB
}

func New(db database.DB) UsersGateway {
	return UsersGateway{db}
}
