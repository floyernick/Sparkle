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

func (service UsersService) Signup(request UsersSignupRequest) (UsersSignupResponse, error) {

	var response UsersSignupResponse

	if err := validator.Process(request); err != nil {
		return response, errors.InvalidParams
	}

	user, err := service.users.GetByUsername(request.Username)

	if err != nil {
		return response, errors.InternalError
	}

	if user.Exists() {
		return response, errors.UsernameUsed
	}

	user = entities.User{
		Username:    request.Username,
		AccessToken: uuid.Generate(),
	}

	user.Password, err = passwords.GenerateHash(request.Password)

	if err != nil {
		return response, errors.InternalError
	}

	user.Id, err = service.users.Create(user)

	if err != nil {
		return response, errors.InternalError
	}

	response = UsersSignupResponse{
		Id:    user.Id,
		Token: user.AccessToken,
	}

	return response, nil

}
