package likes

import (
	"Sparkle/app/errors"
	"Sparkle/entities"
	"Sparkle/tools/validator"
)

type LikesCreateRequest struct {
	Token  string `json:"token" validate:"required,uuid"`
	PostId int    `json:"post_id" validate:"required,min=1"`
}

type LikesCreateResponse struct{}

func (service LikesService) Create(params LikesCreateRequest) (LikesCreateResponse, error) {

	var response LikesCreateResponse

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

	post, err := service.posts.GetById(params.PostId)

	if err != nil {
		return response, errors.InternalError
	}

	if !post.Exists() {
		return response, errors.PostNotFound
	}

	like, err := service.likes.GetByUserIdAndPostId(user.Id, post.Id)

	if err != nil {
		return response, errors.InternalError
	}

	if like.Exists() {
		return response, errors.LikeExists
	}

	like = entities.Like{
		UserId: user.Id,
		PostId: post.Id,
	}

	like.Id, err = service.likes.Create(like)

	if err != nil {
		return response, errors.InternalError
	}

	return response, nil

}
