package delivery

import (
	"net/http"

	"Sparkle/config"
	"Sparkle/delivery/likes"
	"Sparkle/delivery/locations"
	"Sparkle/delivery/posts"
	"Sparkle/delivery/users"
)

func Serve(config config.ServerConfig, users users.UsersController, posts posts.PostsController,
	likes likes.LikesController, locations locations.LocationsController) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/users.signup", users.Signup)
	mux.HandleFunc("/users.signin", users.Signin)
	mux.HandleFunc("/users.get", users.Get)
	mux.HandleFunc("/posts.create", posts.Create)
	mux.HandleFunc("/posts.update", posts.Update)
	mux.HandleFunc("/posts.delete", posts.Delete)
	mux.HandleFunc("/posts.list", posts.List)
	mux.HandleFunc("/likes.create", likes.Create)
	mux.HandleFunc("/likes.delete", likes.Delete)
	mux.HandleFunc("/locations.list", locations.List)

	server := &http.Server{
		Addr:         config.Port,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      http.TimeoutHandler(mux, config.HandlerTimeout, ""),
	}

	return server.ListenAndServe()

}
