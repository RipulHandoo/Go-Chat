package handler

import (
	"net/http"

	"github.com/RipulHandoo/goChat/db/database"
	"github.com/RipulHandoo/goChat/utils"
)

// LogOut is an HTTP handler for user logout.
func LogOut(w http.ResponseWriter, r *http.Request, user database.User) {
	// Clear the authentication token cookie.
	http.SetCookie(w, &http.Cookie{
		Name:   "auth_token",
		Value:  "",
		Path:   "/",
	})

	// Respond with a JSON representation of the user and an "Accepted" status code.
	utils.ResponseWithJson(w, http.StatusAccepted, user)
}
