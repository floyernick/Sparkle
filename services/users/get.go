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

	var result UsersGetResponse

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

	requestedUser, err := service.users.GetById(params.Id)

	if err != nil {
		return result, errors.InternalError
	}

	if !requestedUser.Exists() {
		return result, errors.UserNotFound
	}

	createdAfter := time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339)

	posts, err := service.posts.ListByFilter(posts.ListFilter{
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
