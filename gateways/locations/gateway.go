package locations

import (
	"Sparkle/gateways/common/cache"
	"Sparkle/gateways/common/database"
)

type LocationsGateway struct {
	db    database.DB
	cache cache.Cache
}

func New(db database.DB, cache cache.Cache) LocationsGateway {
	return LocationsGateway{db, cache}
}
