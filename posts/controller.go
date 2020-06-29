package posts

import (
	"Sparkle/database"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db database.DB) {
	mux.HandleFunc("/posts.create", PostsCreateController{db}.Handler)
	mux.HandleFunc("/posts.update", PostsUpdateController{db}.Handler)
}
