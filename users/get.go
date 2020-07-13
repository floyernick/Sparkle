package users

import (
	"Sparkle/app/errors"
	"Sparkle/database"
	"Sparkle/handler"
	"Sparkle/tools/validator"
	"net/http"
	"time"
)

type UsersGetRequest struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    int    `json:"id" validate:"required,min=1"`
}

type UsersGetResponse struct {
	User  UsersGetResponseUser   `json:"user"`
	Posts []UsersGetResponsePost `json:"posts"`
}

type UsersGetResponseUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UsersGetResponsePost struct {
	Id           int    `json:"id"`
	Text         string `json:"text"`
	LocationCode string `json:"location_code"`
	CreatedAt    string `json:"created_at"`
	LikesNumber  int    `json:"likes_number"`
}

type UsersGetController struct {
	db database.DB
}

func (controller UsersGetController) Handler(w http.ResponseWriter, r *http.Request) {

	var req UsersGetRequest

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

func (controller UsersGetController) Usecase(params UsersGetRequest) (UsersGetResponse, error) {

	var result UsersGetResponse

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

	requestedUser, err := controller.db.GetUserById(params.Id)

	if err != nil {
		return result, errors.InternalError
	}

	if !requestedUser.Exists() {
		return result, errors.UserNotFound
	}

	createdAfter := time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339)

	posts, err := controller.db.GetPostsByFilter(database.PostsFilter{
		UserIdEquals:         requestedUser.Id,
		CreatedAfter:         createdAfter,
		OrderByCreatedAtDesc: true,
	})

	if err != nil {
		return result, errors.InternalError
	}

	result.User = UsersGetResponseUser{
		Id:       requestedUser.Id,
		Username: requestedUser.Username,
	}
	result.Posts = make([]UsersGetResponsePost, 0, len(posts))

	for _, post := range posts {
		resultPost := UsersGetResponsePost{
			Id:           post.Id,
			Text:         post.Text,
			LocationCode: post.LocationCode,
			CreatedAt:    post.CreatedAt,
			LikesNumber:  post.LikesNumber,
		}
		result.Posts = append(result.Posts, resultPost)
	}

	return result, nil

}
