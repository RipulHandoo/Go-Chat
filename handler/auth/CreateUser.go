package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RipulHandoo/goChat/db"
	"github.com/RipulHandoo/goChat/db/database"
	"github.com/RipulHandoo/goChat/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	apiConfig := db.DBInstance()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		hashPassword = []byte(params.Password)
	}

	user, err := apiConfig.CreateUser(r.Context(), database.CreateUserParams{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashPassword),
	})

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err)
	}

	token, expiryTime, err := utils.GetJwtToken(utils.Credential{
		Email:    params.Email,
		Username: params.Username,
	})

	if err != nil {
		utils.ResponseWithError(w, http.StatusUnauthorized, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "auth_token",
		Value:   token,
		Expires: expiryTime,
		Path:    "/",
	})

	utils.ResponseWithJson(w, http.StatusAccepted, utils.MapRegisteredUser(user))
}
