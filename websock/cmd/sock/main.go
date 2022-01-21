package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/SarthakJha/distr-websock/internal"
	"github.com/SarthakJha/distr-websock/repository"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err.Error())
	}
	// server config
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")),
		ReadTimeout:  time.Second * 60 * 5,
		WriteTimeout: time.Second * 60 * 5,
		IdleTimeout:  time.Second * 60 * 5,
	}

	msgTable := repository.MessageRepository{}
	usrTable := repository.UserRepository{}
	redisRepo := repository.ConnectionRepository{}
	hler := internal.Chans{}
	hler.InitChan(10)

	redisRepo.InitConnectionRepository()
	msgTable.InitMessageConnection()
	usrTable.InitUserConnection()

	// TODO: make 2 channels
	// TODO: create go routines

	http.HandleFunc("/ws", hler.HandleConn)

	go func() {
		fmt.Println("listening at : ", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	shutSig := make(chan os.Signal, 1)

	signal.Notify(shutSig, os.Interrupt)
	<-shutSig

	// graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer server.Shutdown(ctx)
	defer cancel()
	// TODO: send shutdown signals to all goroutines

}
