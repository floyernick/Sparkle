package likes

import (
	"Sparkle/database"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db database.DB) {
	mux.HandleFunc("/likes.create", LikesCreateController{db}.Handler)
	mux.HandleFunc("/likes.delete", LikesDeleteController{db}.Handler)
}
