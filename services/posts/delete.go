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

func (service PostsService) Delete(request PostsDeleteRequest) (PostsDeleteResponse, error) {

	var response PostsDeleteResponse

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

	post, err := service.posts.GetById(request.Id)

	if err != nil {
		return response, errors.InternalError
	}

	if !post.Exists() {
		return response, errors.PostNotFound
	}

	if !post.CreatedBy(user) {
		return response, errors.ActionNotAllowed
	}

	err = service.posts.Delete(post)

	if err != nil {
		return response, errors.InternalError
	}

	return response, nil

}
