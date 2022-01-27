package internal

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccessTokenCustomClaims struct {
	userID    string
	username  string
	tokenType string
	jwt.StandardClaims
}

func SignUserToken(userID, username, tokenType, issuer string, expiration int) (string, error) {
	claims := AccessTokenCustomClaims{
		userID:    userID,
		username:  username,
		tokenType: "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expiration)).Unix(),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(os.Getenv("JWT_SECRET"))
}
