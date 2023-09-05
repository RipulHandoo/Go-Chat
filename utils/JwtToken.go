package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Credential struct {
	Email    string
	Username string
}

type Claims struct {
	Creds Credential
	jwt.RegisteredClaims
}

func GetJwtToken(signerClaims Credential) (tokenString string, expiry time.Time, Token_err error) {
	godotenv.Load(".env")
	var jwt_Key string = os.Getenv("SECRET_KEY")
	var expiryTime = time.Now().Add(50 * time.Minute)

	claims := Claims{
		Creds: signerClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, Token_err = token.SignedString([]byte(jwt_Key))

	return tokenString, expiryTime, Token_err
}