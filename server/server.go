package server

import (
	"net/http"

	"Sparkle/config"
	"Sparkle/database"
	"Sparkle/posts"
	"Sparkle/users"
)

func Init(config config.ServerConfig, db database.DB) error {

	mux := http.NewServeMux()
	users.RegisterRoutes(mux, db)
	posts.RegisterRoutes(mux, db)

	server := &http.Server{
		Addr:         config.Port,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      mux,
	}

	return server.ListenAndServe()

}
