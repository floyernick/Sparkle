package users

import (
	"Sparkle/app/errors"
	"Sparkle/delivery/common/handlers"
	"Sparkle/services/users"
	"net/http"
)

func (controller UsersController) Signup(w http.ResponseWriter, r *http.Request) {

	var req users.UsersSignupRequest

	if err := handlers.ParseRequestBody(r, &req); err != nil {
		handlers.RespondWithError(w, errors.BadRequest)
	}

	res, err := controller.service.Signup(req)

	if err != nil {
		handlers.RespondWithError(w, err)
	} else {
		handlers.RespondWithSuccess(w, res)
	}

}
