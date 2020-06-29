package posts

import (
	"Sparkle/app/errors"
	"Sparkle/database"
	"Sparkle/handler"
	"Sparkle/tools/validator"
	"net/http"
)

type PostsDeleteRequest struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    int    `json:"id" validate:"required,min=1"`
}

type PostsDeleteResponse struct{}

type PostsDeleteController struct {
	db database.DB
}

func (controller PostsDeleteController) Handler(w http.ResponseWriter, r *http.Request) {

	var req PostsDeleteRequest

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

func (controller PostsDeleteController) Usecase(params PostsDeleteRequest) (PostsDeleteResponse, error) {

	var result PostsDeleteResponse

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

	post, err := controller.db.GetPostById(params.Id)

	if err != nil {
		return result, errors.InternalError
	}

	if !post.Exists() {
		return result, errors.PostNotFound
	}

	if !post.CreatedBy(user) {
		return result, errors.ActionNotAllowed
	}

	err = controller.db.DeletePost(post)

	if err != nil {
		return result, errors.InternalError
	}

	return result, nil

}
