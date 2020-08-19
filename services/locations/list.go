package locations

import (
	"Sparkle/app/errors"
	"Sparkle/entities"
	"Sparkle/gateways/locations"
	"Sparkle/tools/geohash"
	"Sparkle/tools/validator"
	"time"
)

type LocationsListRequest struct {
	Token     string  `json:"token" validate:"required,uuid"`
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180"`
	Latitude  float64 `json:"latitude" validate:"required,min=-90,max=90"`
	Zoom      int     `json:"zoom" validate:"required,min=1,max=20"`
}

type LocationsListResponse struct {
	Locations []LocationsListResponseLocation `json:"locations"`
}

type LocationsListResponseLocation struct {
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	PostsNumber int     `json:"posts_number"`
}

func (service LocationsService) List(request LocationsListRequest) (LocationsListResponse, error) {

	var response LocationsListResponse

	if err := validator.Process(request); err != nil {
		return response, errors.InvalidParams
	}

	user, err := service.users.GetByAccessToken(request.Token)

	if err != nil {
		return response, errors.InternalError
	}

	if !user.Exists() {
		return response, errors.InvalidCredentials
	}

	parentLocationCode := geohash.CoordinatesToGeohash(request.Latitude, request.Longitude, geohash.GetParentLength(request.Zoom))

	childLocationCodeLength := geohash.GetDisplayableLength(request.Zoom)

	createdAfter := time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339)

	locations, err := service.locations.ListByFilter(locations.ListFilter{
		ChildCodeLengthEquals: childLocationCodeLength,
		ParentCodeStartsWith:  parentLocationCode,
		CreatedAfter:          createdAfter,
		Limit:                 100,
	})

	if err != nil {
		return response, errors.InternalError
	}

	locations = entities.LocationsWithCoordinates(locations)

	response.Locations = make([]LocationsListResponseLocation, 0, len(locations))

	for _, location := range locations {
		responseLocation := LocationsListResponseLocation{
			Longitude:   location.Longitude,
			Latitude:    location.Latitude,
			PostsNumber: location.PostsNumber,
		}
		response.Locations = append(response.Locations, responseLocation)
	}

	return response, nil

}
