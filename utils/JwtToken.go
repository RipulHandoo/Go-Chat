package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Credential represents user credentials (email and username).
type Credential struct {
	Email    string
	Username string
}

// Claims represents JWT claims with custom data (Creds) and standard registered claims.
type Claims struct {
	Creds Credential
	jwt.RegisteredClaims
}

// GetJwtToken generates a JWT token with the provided signerClaims and returns the token string,
// expiry time, and any error that occurred during token creation.
func GetJwtToken(signerClaims Credential) (tokenString string, expiry time.Time, Token_err error) {
	// Load environment variables from a .env file.
	godotenv.Load(".env")

	// Get the JWT secret key from environment variables.
	var jwt_Key string = os.Getenv("SECRET_KEY")

	// Define the token's expiration time (e.g., 50 minutes from now).
	var expiryTime = time.Now().Add(50 * time.Minute)

	// Create JWT claims with custom data (signerClaims) and standard registered claims (ExpiresAt).
	claims := Claims{
		Creds: signerClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}

	// Create a new JWT token using HMAC SHA-256 signing method.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the JWT secret key.
	tokenString, Token_err = token.SignedString([]byte(jwt_Key))

	// Return the generated token string, expiry time, and any error.
	return tokenString, expiryTime, Token_err
}
