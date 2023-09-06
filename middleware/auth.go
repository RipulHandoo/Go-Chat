package middleware

import (
	"fmt"
	"net/http"
	// "os"

	"github.com/RipulHandoo/goChat/db"
	"github.com/RipulHandoo/goChat/db/database"
	"github.com/RipulHandoo/goChat/utils"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/joho/godotenv"
)

// AuthHandler is a type for the authentication handler function.
type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

// Auth is a middleware function that checks for user authentication and authorization.
func Auth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Load environment variables from a .env file.
		// godotenv.Load(".env")
		// var jwtString string = os.Getenv("SECRET_KEY")
		var jwtString string = "same for everyone"

		// Get the JWT token from the request's cookies.
		jwtToken, err := r.Cookie("auth_token")

		if err != nil {
			if err == http.ErrNoCookie {
				// If there is no authentication token cookie, respond with unauthorized status.
				fmt.Print("NO COOKIE")
				utils.ResponseWithError(w, http.StatusUnauthorized, err)
				return
			}
			// If there's an error, respond with unauthorized status and the error.
			utils.ResponseWithError(w, http.StatusUnauthorized, err)
			return
		}

		// Get the token string from the cookie.
		tknStr := jwtToken.Value
		claims := &utils.Claims{}

		// Parse the JWT token with claims using the JWT secret key.
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtString), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// If the token signature is invalid, respond with unauthorized status.
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Print("NO JWT TOKEN")
			// If there's an error, respond with unauthorized status and the error.
			utils.ResponseWithError(w, http.StatusUnauthorized, err)
			return
		}

		if !tkn.Valid {
			// If the token is not valid, respond with unauthorized status.
			utils.ResponseWithError(w, http.StatusUnauthorized, err)
			return
		}

		// Extract the user's email from the JWT claims.
		userEmail := claims.Creds.Email
		apiConfig := db.DbClient

		// Retrieve user information from the database based on the email.
		user, err := apiConfig.GetUserByEmail(r.Context(), userEmail)

		if err != nil {
			// If there's an error in retrieving the user, respond with unauthorized status.
			utils.ResponseWithError(w, http.StatusUnauthorized, err)
			return
		}

		// Call the provided authentication handler function with the user information.
		handler(w, r, user)
	}
}
