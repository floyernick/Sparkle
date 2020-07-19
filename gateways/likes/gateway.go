package likes

import (
	"Sparkle/gateways/common/database"
)

type LikesGateway struct {
	db database.DB
}

func New(db database.DB) LikesGateway {
	return LikesGateway{db}
}
