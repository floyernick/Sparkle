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

func (service PostsService) List(request PostsListRequest) (PostsListResponse, error) {

	var response PostsListResponse

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

	locationCode := geohash.CoordinatesToGeohash(request.Latitude, request.Longitude, geohash.GetClickableLength(request.Zoom))

	createdAfter := time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339)

	posts, err := service.posts.ListByFilter(posts.ListFilter{
		LocationCodeStartsWith: locationCode,
		CreatedAfter:           createdAfter,
		OrderByCreatedAtDesc:   true,
		Offset:                 request.Offset,
		Limit:                  request.Limit,
	})

	if err != nil {
		return response, errors.InternalError
	}

	users, err := service.users.ListByIds(entities.GetUserIdsFromPosts(posts))

	if err != nil {
		return response, errors.InternalError
	}

	usersMap := entities.UsersListToMap(users)

	response.Posts = make([]PostsListResponsePost, 0, len(posts))

	for _, post := range posts {
		responsePost := PostsListResponsePost{
			Id:          post.Id,
			Text:        post.Text,
			CreatedAt:   post.CreatedAt,
			LikesNumber: post.LikesNumber,
		}
		user := usersMap[post.UserId]
		responsePost.User = PostsListResponsePostUser{
			Id:       user.Id,
			Username: user.Username,
		}
		response.Posts = append(response.Posts, responsePost)
	}

	return response, nil

}
