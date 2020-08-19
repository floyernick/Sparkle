package users

import (
	"Sparkle/app/errors"
	"Sparkle/gateways/posts"
	"Sparkle/tools/validator"
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

func (service UsersService) Get(params UsersGetRequest) (UsersGetResponse, error) {

	var response UsersGetResponse

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

	requestedUser, err := service.users.GetById(params.Id)

	if err != nil {
		return response, errors.InternalError
	}

	if !requestedUser.Exists() {
		return response, errors.UserNotFound
	}

	createdAfter := time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339)

	posts, err := service.posts.ListByFilter(posts.ListFilter{
		UserIdEquals:         requestedUser.Id,
		CreatedAfter:         createdAfter,
		OrderByCreatedAtDesc: true,
	})

	if err != nil {
		return response, errors.InternalError
	}

	response.User = UsersGetResponseUser{
		Id:       requestedUser.Id,
		Username: requestedUser.Username,
	}
	response.Posts = make([]UsersGetResponsePost, 0, len(posts))

	for _, post := range posts {
		responsePost := UsersGetResponsePost{
			Id:           post.Id,
			Text:         post.Text,
			LocationCode: post.LocationCode,
			CreatedAt:    post.CreatedAt,
			LikesNumber:  post.LikesNumber,
		}
		response.Posts = append(response.Posts, responsePost)
	}

	return response, nil

}
