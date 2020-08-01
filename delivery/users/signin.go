package users

import (
	"Sparkle/app/errors"
	"Sparkle/delivery/common/handlers"
	"Sparkle/services/users"
	"net/http"
)

func (controller UsersController) Signin(w http.ResponseWriter, r *http.Request) {

	var req users.UsersSigninRequest

	if err := handlers.ParseRequestBody(r, &req); err != nil {
		handlers.RespondWithError(w, errors.BadRequest)
		return
	}

	res, err := controller.service.Signin(req)

	if err != nil {
		handlers.RespondWithError(w, err)
		return
	}

	handlers.RespondWithSuccess(w, res)

}
