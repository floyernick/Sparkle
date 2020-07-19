package users

import (
	"Sparkle/services/users"
)

type UsersController struct {
	service users.UsersService
}

func New(service users.UsersService) UsersController {
	return UsersController{service}
}
