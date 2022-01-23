package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/SarthakJha/distr-websock/internal"
	"github.com/SarthakJha/distr-websock/internal/models"
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
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	msgTable := repository.MessageRepository{}
	usrTable := repository.UserRepository{}
	redisRepo := repository.ConnectionRepository{}
	hler := internal.Chans{}
	hler.InitChan(10)

	redisRepo.InitConnectionRepository()
	msgTable.InitMessageConnection()
	usrTable.InitUserConnection()

	chan1 := make(chan models.Message, 10)

	var ctxCan []context.CancelFunc
	wg1.Add(10)
	wg2.Add(10)
	// TODO: create go routines
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		go internal.KafkaPub(*hler.GetPubChan(), &redisRepo, i, &wg1)
		go internal.KafkaSub(chan1, i, ctx)
		ctxCan = append(ctxCan, cancel)
		go internal.WSWriterHandler(chan1, hler.GetMap(), msgTable, &wg2)
	}

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
	// shutting down publishing workers
	for _, val := range ctxCan {
		// sending cancel signal to all the producing workers
		val()
	}
	server.Shutdown(ctx)
	wg1.Wait()
	wg2.Wait()
	defer cancel()
	// TODO: send shutdown signals to all goroutines

}
