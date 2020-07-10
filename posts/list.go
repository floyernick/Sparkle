package posts

import (
	"Sparkle/app/errors"
	"Sparkle/database"
	"Sparkle/entities"
	"Sparkle/handler"
	"Sparkle/tools/geohash"
	"Sparkle/tools/validator"
	"net/http"
	"time"
)

type PostsListRequest struct {
	Token     string  `json:"token" validate:"required,uuid"`
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180"`
	Latitude  float64 `json:"latitude" validate:"min=-90,max=90"`
	Zoom      int     `json:"zoom" validate:"min=1,max=20"`
	Offset    int     `json:"offset" validate:"min=0"`
	Limit     int     `json:"limit" validate:"min=1,max=50"`
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

type PostsListController struct {
	db database.DB
}

func (controller PostsListController) Handler(w http.ResponseWriter, r *http.Request) {

	var req PostsListRequest

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

func (controller PostsListController) Usecase(params PostsListRequest) (PostsListResponse, error) {

	var result PostsListResponse

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

	locationCode := geohash.CoordinatesToGeohash(params.Latitude, params.Longitude, geohash.GetClickableLength(params.Zoom))

	createdAfter := time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339)

	posts, err := controller.db.GetPostsByFilter(database.PostsFilter{
		LocationCodeStartsWith: locationCode,
		CreatedAfter:           createdAfter,
		OrderByIdDesc:          true,
		Offset:                 params.Offset,
		Limit:                  params.Limit,
	})

	if err != nil {
		return result, errors.InternalError
	}

	users, err := controller.db.GetUsersByIds(entities.GetUserIdsFromPosts(posts))

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
