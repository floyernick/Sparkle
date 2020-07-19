package posts

import (
	"Sparkle/app/errors"
	"Sparkle/tools/validator"
)

type PostsDeleteRequest struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    int    `json:"id" validate:"required,min=1"`
}

type PostsDeleteResponse struct{}

func (service PostsService) Delete(params PostsDeleteRequest) (PostsDeleteResponse, error) {

	var result PostsDeleteResponse

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

	err = service.posts.Delete(post)

	if err != nil {
		return result, errors.InternalError
	}

	return result, nil

}
