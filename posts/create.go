package posts

import (
	"Sparkle/app/errors"
	"Sparkle/database"
	"Sparkle/entities"
	"Sparkle/handler"
	"Sparkle/tools/locations"
	"Sparkle/tools/validator"
	"net/http"
	"time"
)

type PostsCreateRequest struct {
	Token     string  `json:"token" validate:"required,uuid"`
	Text      string  `json:"text" validate:"required,min=1,max=150"`
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180"`
	Latitude  float64 `json:"latitude" validate:"min=-90,max=90"`
}

type PostsCreateResponse struct {
	Id int `json:"id"`
}

type PostsCreateController struct {
	db database.DB
}

func (controller PostsCreateController) Handler(w http.ResponseWriter, r *http.Request) {

	var req PostsCreateRequest

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

func (controller PostsCreateController) Usecase(params PostsCreateRequest) (PostsCreateResponse, error) {

	var result PostsCreateResponse

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

	post := entities.Post{
		UserId:       user.Id,
		Text:         params.Text,
		LocationCode: locations.ConvertToOLC(params.Latitude, params.Longitude),
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
	}

	post.Id, err = controller.db.CreatePost(post)

	if err != nil {
		return result, errors.InternalError
	}

	result = PostsCreateResponse{
		Id: post.Id,
	}

	return result, nil

}
