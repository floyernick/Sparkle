package posts

import (
	"Sparkle/gateways/posts"
	"Sparkle/gateways/users"
)

type PostsService struct {
	users users.UsersGateway
	posts posts.PostsGateway
}

func New(users users.UsersGateway, posts posts.PostsGateway) PostsService {
	return PostsService{users, posts}
}
