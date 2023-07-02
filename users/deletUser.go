package users

import (
	"net/http"

	"github.com/RipulHandoo/goChat/db"
	"github.com/RipulHandoo/goChat/db/database"
	"github.com/RipulHandoo/goChat/utils"
)

func DeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	apiConfig := db.DBInstance()
	user_id := user.ID

	http.SetCookie(w, &http.Cookie{
		Name:  "auth_token",
		Value: "",
		Path:  "/",
	})

	user, err := apiConfig.DeleteUser(r.Context(), user_id)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, utils.MapDeleteUser(user))
}
