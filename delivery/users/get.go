package users

import (
	"Sparkle/app/errors"
	"Sparkle/delivery/common/handlers"
	"Sparkle/services/users"
	"net/http"
)

func (controller UsersController) Get(w http.ResponseWriter, r *http.Request) {

	var req users.UsersGetRequest

	if err := handlers.ParseRequestBody(r, &req); err != nil {
		handlers.RespondWithError(w, errors.BadRequest)
	}

	res, err := controller.service.Get(req)

	if err != nil {
		handlers.RespondWithError(w, err)
	} else {
		handlers.RespondWithSuccess(w, res)
	}

}
