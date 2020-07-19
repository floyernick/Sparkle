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

func (service UsersService) Signin(params UsersSigninRequest) (UsersSigninResponse, error) {

	var result UsersSigninResponse

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams
	}

	user, err := service.users.GetByUsername(params.Username)

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
