package posts

import (
	"Sparkle/app/errors"
	"Sparkle/database"
	"Sparkle/handler"
	"Sparkle/tools/validator"
	"net/http"
)

type PostsUpdateRequest struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    int    `json:"id" validate:"required,min=1"`
	Text  string `json:"text" validate:"required,min=1,max=150"`
}

type PostsUpdateResponse struct{}

type PostsUpdateController struct {
	db database.DB
}

func (controller PostsUpdateController) Handler(w http.ResponseWriter, r *http.Request) {

	var req PostsUpdateRequest

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

func (controller PostsUpdateController) Usecase(params PostsUpdateRequest) (PostsUpdateResponse, error) {

	var result PostsUpdateResponse

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

	post.Text = params.Text

	err = controller.db.UpdatePost(post)

	if err != nil {
		return result, errors.InternalError
	}

	return result, nil

}
