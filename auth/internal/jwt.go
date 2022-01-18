package internal

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccessTokenCustomClaims struct {
	userID    string
	username  string
	tokenType string
	jwt.StandardClaims
}

func SignUserToken(userID, username, tokenType string) (string, error) {
	expiration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	claims := AccessTokenCustomClaims{
		userID:    userID,
		username:  username,
		tokenType: "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expiration)).Unix(),
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(os.Getenv("JWT_SECRET"))
}
