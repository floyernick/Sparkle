package users

import (
	"Sparkle/app/errors"
	"Sparkle/database"
	"Sparkle/entities"
	"Sparkle/handler"
	"Sparkle/tools/passwords"
	"Sparkle/tools/uuid"
	"Sparkle/tools/validator"
	"net/http"
)

type UsersSignupRequest struct {
	Username string `json:"username" validate:"required,min=1,max=25"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type UsersSignupResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type UsersSignupController struct {
	db database.DB
}

func (controller UsersSignupController) Handler(w http.ResponseWriter, r *http.Request) {

	var req UsersSignupRequest

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

func (controller UsersSignupController) Usecase(params UsersSignupRequest) (UsersSignupResponse, error) {

	var result UsersSignupResponse

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams
	}

	user, err := controller.db.GetUserByUsername(params.Username)

	if err != nil {
		return result, errors.InternalError
	}

	if user.Exists() {
		return result, errors.UsernameUsed
	}

	user = entities.User{
		Username:    params.Username,
		AccessToken: uuid.Generate(),
	}

	user.Password, err = passwords.GenerateHash(params.Password)

	if err != nil {
		return result, errors.InternalError
	}

	user.Id, err = controller.db.CreateUser(user)

	if err != nil {
		return result, errors.InternalError
	}

	result = UsersSignupResponse{
		Id:    user.Id,
		Token: user.AccessToken,
	}

	return result, nil

}
