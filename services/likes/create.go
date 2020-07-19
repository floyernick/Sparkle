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

	var result LikesCreateResponse

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

	post, err := service.posts.GetById(params.PostId)

	if err != nil {
		return result, errors.InternalError
	}

	if !post.Exists() {
		return result, errors.PostNotFound
	}

	like, err := service.likes.GetByUserIdAndPostId(user.Id, post.Id)

	if err != nil {
		return result, errors.InternalError
	}

	if like.Exists() {
		return result, errors.LikeExists
	}

	like = entities.Like{
		UserId: user.Id,
		PostId: post.Id,
	}

	like.Id, err = service.likes.Create(like)

	if err != nil {
		return result, errors.InternalError
	}

	return result, nil

}
