package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RipulHandoo/goChat/db"
	"github.com/RipulHandoo/goChat/utils"
	"golang.org/x/crypto/bcrypt"
)


func LoginUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Email    string	`json:"email"`
		Password string		`json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil{
		utils.ResponseWithError(w,http.StatusBadRequest,err)
		return
	}

	apiConfig := db.DbClient
	user, err := apiConfig.GetUserByEmail(r.Context(),params.Email)

	if err != nil{
		utils.ResponseWithError(w,401,err)
	}

	authCheck := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(params.Password))
	if authCheck != nil{
		utils.ResponseWithError(w,http.StatusUnauthorized,authCheck)
		return
	}

	token, expiry, err := utils.GetJwtToken(utils.Credential{
		Email: user.Email,
		Username: user.Username,
	})

	if err != nil{
		utils.ResponseWithError(w,http.StatusForbidden,err)
		return
	}

	http.SetCookie(w,&http.Cookie{
		Name: "auth_token",
		Value: token,
		Expires: expiry,
		Path: "/",
	})

	utils.ResponseWithJson(w,200,utils.MapLoginUser(user))
}