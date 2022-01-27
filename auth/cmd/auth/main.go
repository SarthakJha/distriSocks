package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/SarthakJha/distrisock-auth/internal/api"
	"github.com/SarthakJha/distrisock-auth/internal/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err.Error())
	}
	conf, err := utils.LoadConfig("../../config/config.prod.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	r := mux.NewRouter()
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", conf.PORT),
		ReadTimeout:  time.Second * 60 * 5,
		WriteTimeout: time.Second * 60 * 5,
		IdleTimeout:  time.Second * 60 * 5,
		Handler:      r,
	}

	api.SetupRoutes(r)

	go func() {
		fmt.Println("listening at : ", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt)
	<-sig

	// graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer server.Shutdown(ctx)
	defer cancel()
	fmt.Println("shutting down...")
}
