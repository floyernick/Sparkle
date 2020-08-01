package posts

import (
	"Sparkle/app/errors"
	"Sparkle/delivery/common/handlers"
	"Sparkle/services/posts"
	"net/http"
)

func (controller PostsController) Create(w http.ResponseWriter, r *http.Request) {

	var req posts.PostsCreateRequest

	if err := handlers.ParseRequestBody(r, &req); err != nil {
		handlers.RespondWithError(w, errors.BadRequest)
		return
	}

	res, err := controller.service.Create(req)

	if err != nil {
		handlers.RespondWithError(w, err)
		return
	}

	handlers.RespondWithSuccess(w, res)

}
