package posts

import (
	"Sparkle/app/errors"
	"Sparkle/entities"
	"Sparkle/gateways/posts"
	"Sparkle/tools/geohash"
	"Sparkle/tools/validator"
	"time"
)

type PostsListRequest struct {
	Token     string  `json:"token" validate:"required,uuid"`
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180"`
	Latitude  float64 `json:"latitude" validate:"required,min=-90,max=90"`
	Zoom      int     `json:"zoom" validate:"required,min=1,max=20"`
	Offset    int     `json:"offset" validate:"required,min=0"`
	Limit     int     `json:"limit" validate:"required,min=1,max=100"`
}

type PostsListResponse struct {
	Posts []PostsListResponsePost `json:"posts"`
}

type PostsListResponsePost struct {
	Id          int                       `json:"id"`
	Text        string                    `json:"text"`
	CreatedAt   string                    `json:"created_at"`
	LikesNumber int                       `json:"likes_number"`
	User        PostsListResponsePostUser `json:"user"`
}

type PostsListResponsePostUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (service PostsService) List(params PostsListRequest) (PostsListResponse, error) {

	var result PostsListResponse

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

	locationCode := geohash.CoordinatesToGeohash(params.Latitude, params.Longitude, geohash.GetClickableLength(params.Zoom))

	createdAfter := time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339)

	posts, err := service.posts.ListByFilter(posts.ListFilter{
		LocationCodeStartsWith: locationCode,
		CreatedAfter:           createdAfter,
		OrderByCreatedAtDesc:   true,
		Offset:                 params.Offset,
		Limit:                  params.Limit,
	})

	if err != nil {
		return result, errors.InternalError
	}

	users, err := service.users.ListByIds(entities.GetUserIdsFromPosts(posts))

	if err != nil {
		return result, errors.InternalError
	}

	usersMap := entities.UsersListToMap(users)

	result.Posts = make([]PostsListResponsePost, 0, len(posts))

	for _, post := range posts {
		resultPost := PostsListResponsePost{
			Id:          post.Id,
			Text:        post.Text,
			CreatedAt:   post.CreatedAt,
			LikesNumber: post.LikesNumber,
		}
		user := usersMap[post.UserId]
		resultPost.User = PostsListResponsePostUser{
			Id:       user.Id,
			Username: user.Username,
		}
		result.Posts = append(result.Posts, resultPost)
	}

	return result, nil

}
