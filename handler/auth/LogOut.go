package handler

import (
	"net/http"

	"github.com/RipulHandoo/goChat/db/database"
	"github.com/RipulHandoo/goChat/utils"
)

func LogOut(w http.ResponseWriter, r *http.Request, user database.User){
	// clear cookies
	
	http.SetCookie(w,&http.Cookie{
		Name: "auth_token",
		Value:"",
		Path: "/",
	})

	utils.ResponseWithJson(w,http.StatusAccepted,user)
}