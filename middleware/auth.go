package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RipulHandoo/goChat/db"
	"github.com/RipulHandoo/goChat/db/database"
	"github.com/RipulHandoo/goChat/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type AuthHandler func(http.ResponseWriter,*http.Request, database.User)

func Auth(handler AuthHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		godotenv.Load(".env")
		var jwtString string = os.Getenv("SECRET_KEY")
		jwtToken, err := r.Cookie("auth_token")

		if err != nil{
			if err == http.ErrNoCookie{
				fmt.Print("NO COOKIE")
				utils.ResponseWithError(w,http.StatusUnauthorized,err)
				return
			}
			utils.ResponseWithError(w,http.StatusUnauthorized,err)
			return
		}

		tknStr := jwtToken.Value
		claims := &utils.Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr,claims,func(token *jwt.Token)(interface{},error){
			return []byte(jwtString), nil
		})
		if err != nil{
			if err == jwt.ErrSignatureInvalid{
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Print("NO JWT TOKEN")
			utils.ResponseWithError(w,http.StatusUnauthorized,err)
			return
		}

		if !tkn.Valid {
			utils.ResponseWithError(w,http.StatusUnauthorized,err)
			return
		}

		userEmail := claims.Creds.Email
		apiConfig := db.DbClient

		user, err := apiConfig.GetUserByEmail(r.Context(),userEmail)

		if err != nil{
			utils.ResponseWithError(w,http.StatusUnauthorized,err)
			return
		}

		handler(w,r,user)
	}
}