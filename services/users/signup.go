package users

import (
	"Sparkle/app/errors"
	"Sparkle/entities"
	"Sparkle/tools/passwords"
	"Sparkle/tools/uuid"
	"Sparkle/tools/validator"
)

type UsersSignupRequest struct {
	Username string `json:"username" validate:"required,min=1,max=25"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type UsersSignupResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func (service UsersService) Signup(params UsersSignupRequest) (UsersSignupResponse, error) {

	var result UsersSignupResponse

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams
	}

	user, err := service.users.GetByUsername(params.Username)

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

	user.Id, err = service.users.Create(user)

	if err != nil {
		return result, errors.InternalError
	}

	result = UsersSignupResponse{
		Id:    user.Id,
		Token: user.AccessToken,
	}

	return result, nil

}
