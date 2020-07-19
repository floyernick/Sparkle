package posts

import (
	"Sparkle/services/posts"
)

type PostsController struct {
	service posts.PostsService
}

func New(service posts.PostsService) PostsController {
	return PostsController{service}
}
