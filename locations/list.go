package locations

import (
	"Sparkle/app/errors"
	"Sparkle/database"
	"Sparkle/entities"
	"Sparkle/handler"
	"Sparkle/tools/geohash"
	"Sparkle/tools/validator"
	"net/http"
	"time"
)

type LocationsListRequest struct {
	Token     string  `json:"token" validate:"required,uuid"`
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180"`
	Latitude  float64 `json:"latitude" validate:"min=-90,max=90"`
	Zoom      int     `json:"zoom" validate:"min=1,max=20"`
}

type LocationsListResponse struct {
	Locations []LocationsListResponseLocation `json:"locations"`
}

type LocationsListResponseLocation struct {
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	PostsNumber int     `json:"posts_number"`
}

type LocationsListController struct {
	db database.DB
}

func (controller LocationsListController) Handler(w http.ResponseWriter, r *http.Request) {

	var req LocationsListRequest

	if err := handler.ParseRequestBody(r, &req); err != nil {
		handler.RespondWithError(w, errors.BadRequest)
	}

	res, err := controller.Usecase(req)

	if err != nil {
		handler.RespondWithError(w, err)
	} else {
		handler.RespondWithSuccess(w, res)
	}

}

func (controller LocationsListController) Usecase(params LocationsListRequest) (LocationsListResponse, error) {

	var result LocationsListResponse

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams
	}

	user, err := controller.db.GetUserByAccessToken(params.Token)

	if err != nil {
		return result, errors.InternalError
	}

	if !user.Exists() {
		return result, errors.InvalidCredentials
	}

	parentLocationCode := geohash.CoordinatesToGeohash(params.Latitude, params.Longitude, geohash.GetParentLength(params.Zoom))

	childLocationCodeLength := geohash.GetDisplayableLength(params.Zoom)

	createdAfter := time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339)

	locations, err := controller.db.GetLocationsByFilter(database.LocationsFilter{
		ChildCodeLengthEquals: childLocationCodeLength,
		ParentCodeStartsWith:  parentLocationCode,
		CreatedAfter:          createdAfter,
		Limit:                 500,
	})

	if err != nil {
		return result, errors.InternalError
	}

	locations = entities.LocationsWithCoordinates(locations)

	for _, location := range locations {
		resultLocation := LocationsListResponseLocation{
			Longitude:   location.Longitude,
			Latitude:    location.Latitude,
			PostsNumber: location.PostsNumber,
		}
		result.Locations = append(result.Locations, resultLocation)
	}

	return result, nil

}
