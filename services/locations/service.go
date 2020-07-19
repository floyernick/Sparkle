package locations

import (
	"Sparkle/gateways/locations"
	"Sparkle/gateways/users"
)

type LocationsService struct {
	users     users.UsersGateway
	locations locations.LocationsGateway
}

func New(users users.UsersGateway, locations locations.LocationsGateway) LocationsService {
	return LocationsService{users, locations}
}
