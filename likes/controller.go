package likes

import (
	"Sparkle/database"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db database.DB) {
	mux.HandleFunc("/likes.create", LikesCreateController{db}.Handler)
}
