package users

import (
	"Sparkle/gateways/posts"
	"Sparkle/gateways/users"
)

type UsersService struct {
	users users.UsersGateway
	posts posts.PostsGateway
}

func New(users users.UsersGateway, posts posts.PostsGateway) UsersService {
	return UsersService{users, posts}
}
