package posts

import (
	"Sparkle/app/errors"
	"Sparkle/entities"
	"Sparkle/tools/geohash"
	"Sparkle/tools/validator"
	"time"
)

type PostsCreateRequest struct {
	Token     string  `json:"token" validate:"required,uuid"`
	Text      string  `json:"text" validate:"required,min=1,max=150"`
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180"`
	Latitude  float64 `json:"latitude" validate:"required,min=-90,max=90"`
}

type PostsCreateResponse struct {
	Id int `json:"id"`
}

func (service PostsService) Create(params PostsCreateRequest) (PostsCreateResponse, error) {

	var result PostsCreateResponse

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams
	}

	user, err := service.users.GetByAccessToken(params.Token)

	if err != nil {
		return result, errors.InternalError
	}

	if !user.Exists() {
		return result, errors.InvalidCredentials
	}

	post := entities.Post{
		UserId:       user.Id,
		Text:         params.Text,
		LocationCode: geohash.CoordinatesToGeohash(params.Latitude, params.Longitude, geohash.MAX_LENGTH),
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
	}

	post.Id, err = service.posts.Create(post)

	if err != nil {
		return result, errors.InternalError
	}

	result = PostsCreateResponse{
		Id: post.Id,
	}

	return result, nil

}
