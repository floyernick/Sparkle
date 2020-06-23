package users

import (
	"Sparkle/database"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db database.DB) {
	mux.HandleFunc("/users.signup", UsersSignupController{db}.Handler)
	mux.HandleFunc("/users.signin", UsersSigninController{db}.Handler)
}
