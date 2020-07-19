package likes

import (
	"Sparkle/gateways/likes"
	"Sparkle/gateways/posts"
	"Sparkle/gateways/users"
)

type LikesService struct {
	users users.UsersGateway
	posts posts.PostsGateway
	likes likes.LikesGateway
}

func New(users users.UsersGateway, posts posts.PostsGateway, likes likes.LikesGateway) LikesService {
	return LikesService{users, posts, likes}
}
