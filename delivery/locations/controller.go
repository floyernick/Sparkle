package locations

import "Sparkle/services/locations"

type LocationsController struct {
	service locations.LocationsService
}

func New(service locations.LocationsService) LocationsController {
	return LocationsController{service}
}
