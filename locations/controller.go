package locations

import (
	"Sparkle/database"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db database.DB) {
	mux.HandleFunc("/locations.list", LocationsListController{db}.Handler)
}
