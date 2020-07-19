package likes

import "Sparkle/services/likes"

type LikesController struct {
	service likes.LikesService
}

func New(service likes.LikesService) LikesController {
	return LikesController{service}
}
