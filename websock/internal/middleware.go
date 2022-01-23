package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/SarthakJha/distr-websock/internal/models"
	"github.com/SarthakJha/distr-websock/internal/utils"
)

type UserData struct {
	Username string
	UserID   string
}

func UpgradeValidation(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-type", "application/json")
		fmt.Println("validating access token")

		token, err := extractToken(r)
		if err != nil {
			utils.ToJSON(models.GenericResponse{
				Success: false,
				Data: models.ErrorResponse{
					Error: err.Error(),
				},
			}, rw)
			return
		}

		userID, username, err := ValidateAccessToken(token)
		if err != nil {
			utils.ToJSON(models.GenericResponse{
				Success: false,
				Data: models.ErrorResponse{
					Error: err.Error(),
				},
			}, rw)
			return
		}
		fmt.Println("validation complete, token correct")

		// passing values to next handler
		ctx := context.WithValue(context.Background(), UserData{}, UserData{
			UserID:   userID,
			Username: username,
		})
		r = r.WithContext(ctx)

		h.ServeHTTP(rw, r)
	})
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")

	if len(authHeaderContent) != 2 {
		return "", errors.New("cant parse auth header")
	}
	return authHeaderContent[1], nil
}
