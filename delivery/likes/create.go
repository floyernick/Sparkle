package likes

import (
	"Sparkle/app/errors"
	"Sparkle/delivery/common/handlers"
	"Sparkle/services/likes"
	"net/http"
)

func (controller LikesController) Create(w http.ResponseWriter, r *http.Request) {

	var req likes.LikesCreateRequest

	if err := handlers.ParseRequestBody(r, &req); err != nil {
		handlers.RespondWithError(w, errors.BadRequest)
	}

	res, err := controller.service.Create(req)

	if err != nil {
		handlers.RespondWithError(w, err)
	} else {
		handlers.RespondWithSuccess(w, res)
	}

}
