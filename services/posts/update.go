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

	var result PostsUpdateResponse

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

	post, err := service.posts.GetById(params.Id)

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

	err = service.posts.Update(post)

	if err != nil {
		return result, errors.InternalError
	}

	return result, nil

}
