package users

import (
	"Sparkle/app/errors"
	"Sparkle/tools/passwords"
	"Sparkle/tools/validator"
)

type UsersSigninRequest struct {
	Username string `json:"username" validate:"required,min=1,max=25"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type UsersSigninResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func (service UsersService) Signin(request UsersSigninRequest) (UsersSigninResponse, error) {

	var response UsersSigninResponse

	if err := validator.Process(request); err != nil {
		return response, errors.InvalidParams
	}

	user, err := service.users.GetByUsername(request.Username)

	if err != nil {
		return response, errors.InternalError
	}

	if !user.Exists() {
		return response, errors.InvalidCredentials
	}

	if !passwords.CheckHash(request.Password, user.Password) {
		return response, errors.InvalidCredentials
	}

	response = UsersSigninResponse{
		Id:    user.Id,
		Token: user.AccessToken,
	}

	return response, nil

}
