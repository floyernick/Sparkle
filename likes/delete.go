package likes

import (
	"Sparkle/app/errors"
	"Sparkle/database"
	"Sparkle/handler"
	"Sparkle/tools/validator"
	"net/http"
)

type LikesDeleteRequest struct {
	Token  string `json:"token" validate:"required,uuid"`
	PostId int    `json:"post_id" validate:"required,min=1"`
}

type LikesDeleteResponse struct{}

type LikesDeleteController struct {
	db database.DB
}

func (controller LikesDeleteController) Handler(w http.ResponseWriter, r *http.Request) {

	var req LikesDeleteRequest

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

func (controller LikesDeleteController) Usecase(params LikesDeleteRequest) (LikesDeleteResponse, error) {

	var result LikesDeleteResponse

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

	post, err := controller.db.GetPostById(params.PostId)

	if err != nil {
		return result, errors.InternalError
	}

	if !post.Exists() {
		return result, errors.PostNotFound
	}

	like, err := controller.db.GetLikeByUserIdAndPostId(user.Id, post.Id)

	if err != nil {
		return result, errors.InternalError
	}

	if !like.Exists() {
		return result, errors.LikeNotFound
	}

	err = controller.db.DeleteLike(like)

	if err != nil {
		return result, errors.InternalError
	}

	return result, nil

}
