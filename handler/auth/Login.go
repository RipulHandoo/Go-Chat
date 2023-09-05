package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RipulHandoo/goChat/db"
	"github.com/RipulHandoo/goChat/utils"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser is an HTTP handler for user login.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Define a structure to hold request parameters.
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Create a JSON decoder for reading the request body.
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	// Decode the JSON request body into the 'params' structure.
	err := decoder.Decode(&params)
	if err != nil {
		// If there is an error in decoding, respond with a bad request error.
		utils.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	// Get a database API instance.
	apiConfig := db.DbClient

	// Retrieve user information from the database based on the provided email.
	user, err := apiConfig.GetUserByEmail(r.Context(), params.Email)

	if err != nil {
		// If there is an error in retrieving the user, respond with an unauthorized error.
		utils.ResponseWithError(w, http.StatusUnauthorized, err)
		return
	}

	// Compare the provided password with the stored hashed password.
	authCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if authCheck != nil {
		// If the password comparison fails, respond with an unauthorized error.
		utils.ResponseWithError(w, http.StatusUnauthorized, authCheck)
		return
	}

	// Generate a JWT token for the authenticated user.
	token, expiry, err := utils.GetJwtToken(utils.Credential{
		Email:    user.Email,
		Username: user.Username,
	})

	if err != nil {
		// If there is an error in token generation, respond with a forbidden error.
		utils.ResponseWithError(w, http.StatusForbidden, err)
		return
	}

	// Set the JWT token as a cookie in the HTTP response.
	http.SetCookie(w, &http.Cookie{
		Name:    "auth_token",
		Value:   token,
		Expires: expiry,
		Path:    "/",
	})

	// Respond with a JSON representation of the authenticated user.
	utils.ResponseWithJson(w, http.StatusOK, utils.MapLoginUser(user))
}
