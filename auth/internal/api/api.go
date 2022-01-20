package api

import (
	"github.com/SarthakJha/distrisock-auth/internal/handler"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", handler.Greet)
}
