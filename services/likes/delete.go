package likes

import (
	"Sparkle/app/errors"
	"Sparkle/tools/validator"
)

type LikesDeleteRequest struct {
	Token  string `json:"token" validate:"required,uuid"`
	PostId int    `json:"post_id" validate:"required,min=1"`
}

type LikesDeleteResponse struct{}

func (service LikesService) Delete(params LikesDeleteRequest) (LikesDeleteResponse, error) {

	var result LikesDeleteResponse

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

	if !like.Exists() {
		return result, errors.LikeNotFound
	}

	err = service.likes.Delete(like)

	if err != nil {
		return result, errors.InternalError
	}

	return result, nil

}
