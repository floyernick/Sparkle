package posts

import (
	"Sparkle/app/errors"
	"Sparkle/tools/validator"
)

type PostsUpdateRequest struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    int    `json:"id" validate:"required,min=1"`
	Text  string `json:"text" validate:"required,min=1,max=150"`
}

type PostsUpdateResponse struct{}

func (service PostsService) Update(params PostsUpdateRequest) (PostsUpdateResponse, error) {

	var response PostsUpdateResponse

	if err := validator.Process(params); err != nil {
		return response, errors.InvalidParams
	}

	user, err := service.users.GetByAccessToken(params.Token)

	if err != nil {
		return response, errors.InternalError
	}

	if !user.Exists() {
		return response, errors.InvalidCredentials
	}

	post, err := service.posts.GetById(params.Id)

	if err != nil {
		return response, errors.InternalError
	}

	if !post.Exists() {
		return response, errors.PostNotFound
	}

	if !post.CreatedBy(user) {
		return response, errors.ActionNotAllowed
	}

	post.Text = params.Text

	err = service.posts.Update(post)

	if err != nil {
		return response, errors.InternalError
	}

	return response, nil

}
