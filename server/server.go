package server

import (
	"net/http"

	"Sparkle/config"
	"Sparkle/database"
	"Sparkle/notes"
)

func Sparkle(config config.ServerConfig, db database.DB) error {

	mux := http.NewServeMux()
	notes.RegisterRoutes(mux, db)

	server := &http.Server{
		Addr:         config.Port,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      mux,
	}

	return server.ListenAndServe()

}
