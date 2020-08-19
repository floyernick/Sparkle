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

func (service LikesService) Delete(request LikesDeleteRequest) (LikesDeleteResponse, error) {

	var response LikesDeleteResponse

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

	post, err := service.posts.GetById(request.PostId)

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

	if !like.Exists() {
		return response, errors.LikeNotFound
	}

	err = service.likes.Delete(like)

	if err != nil {
		return response, errors.InternalError
	}

	return response, nil

}
