package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RipulHandoo/goChat/db"
	"github.com/RipulHandoo/goChat/db/database"
	"github.com/RipulHandoo/goChat/utils"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is an HTTP handler for creating a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Define a structure to hold request parameters.
	type parameters struct {
		Username string `json:"username"`
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
	apiConfig := db.DBInstance()

	// Hash the user's password before storing it in the database.
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		// If there is an error in hashing, use the original password as a fallback.
		hashPassword = []byte(params.Password)
	}

	// Create a new user in the database.
	user, err := apiConfig.CreateUser(r.Context(), database.CreateUserParams{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashPassword),
	})

	if err != nil {
		// If there is an error in creating the user, respond with a bad request error.
		utils.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	// Generate a JWT token for the newly created user.
	token, expiryTime, err := utils.GetJwtToken(utils.Credential{
		Email:    params.Email,
		Username: params.Username,
	})

	if err != nil {
		// If there is an error in token generation, respond with an unauthorized error.
		utils.ResponseWithError(w, http.StatusUnauthorized, err)
		return
	}

	// Set the JWT token as a cookie in the HTTP response.
	http.SetCookie(w, &http.Cookie{
		Name:    "auth_token",
		Value:   token,
		Expires: expiryTime,
		Path:    "/",
	})

	// Respond with a JSON representation of the newly registered user.
	utils.ResponseWithJson(w, http.StatusCreated, utils.MapRegisteredUser(user))
}
