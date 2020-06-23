package users

import (
	"Sparkle/app/errors"
	"Sparkle/database"
	"Sparkle/handler"
	"Sparkle/tools/passwords"
	"Sparkle/tools/validator"
	"net/http"
)

type UsersSigninRequest struct {
	Username string `json:"username" validate:"required,min=1,max=25"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type UsersSigninResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type UsersSigninController struct {
	db database.DB
}

func (controller UsersSigninController) Handler(w http.ResponseWriter, r *http.Request) {

	var req UsersSigninRequest

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

func (controller UsersSigninController) Usecase(params UsersSigninRequest) (UsersSigninResponse, error) {

	var result UsersSigninResponse

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams
	}

	user, err := controller.db.GetUserByUsername(params.Username)

	if err != nil {
		return result, errors.InternalError
	}

	if !user.Exists() {
		return result, errors.InvalidCredentials
	}

	if !passwords.CheckHash(params.Password, user.Password) {
		return result, errors.InvalidCredentials
	}

	result = UsersSigninResponse{
		Id:    user.Id,
		Token: user.AccessToken,
	}

	return result, nil

}
