package posts

import (
	"Sparkle/app/errors"
	"Sparkle/delivery/common/handlers"
	"Sparkle/services/posts"
	"net/http"
)

func (controller PostsController) Update(w http.ResponseWriter, r *http.Request) {

	var req posts.PostsUpdateRequest

	if err := handlers.ParseRequestBody(r, &req); err != nil {
		handlers.RespondWithError(w, errors.BadRequest)
		return
	}

	res, err := controller.service.Update(req)

	if err != nil {
		handlers.RespondWithError(w, err)
		return
	}

	handlers.RespondWithSuccess(w, res)

}
