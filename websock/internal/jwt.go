package internal

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type AccessTokenCustomClaims struct {
	userID    string
	username  string
	tokenType string
	jwt.StandardClaims
}

func ValidateAccessToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, AccessTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			fmt.Println("different siging method")
			err := errors.New("different siging method")
			return nil, err
		}

		return os.Getenv("JWT_SECRET"), nil
	})
	if err != nil {
		return "", "", errors.New("cannot parse token")
	}
	claims := token.Claims.(*AccessTokenCustomClaims)

	if !token.Valid || claims.userID == "" || claims.username == "" {
		return "", "", errors.New("invalid token")
	}

	return claims.userID, claims.username, nil
}
