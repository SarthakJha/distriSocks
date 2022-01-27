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
	"github.com/SarthakJha/distr-websock/internal/utils"
	"github.com/SarthakJha/distr-websock/repository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
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
	// server config
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")),
		ReadTimeout:  time.Second * 60 * 5,
		WriteTimeout: time.Second * 60 * 5,
		IdleTimeout:  time.Second * 60 * 5,
		Handler:      r,
	}
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	msgTable := repository.MessageRepository{}
	usrTable := repository.UserRepository{}
	redisRepo := repository.ConnectionRepository{}
	hler := internal.Chans{}
	hler.InitChan(10)

	redisRepo.InitConnectionRepository(conf.REDIS_SERVICE, conf.REDIS_PORT)
	msgTable.InitMessageConnection(conf.AWS_REGION, conf.MESSAGE_TABLE_NAME)
	usrTable.InitUserConnection(conf.AWS_REGION, conf.USER_TABLE_NAME)

	chan1 := make(chan models.Message, 10)
	kafkaSubCtx, kafkaSubCancel := context.WithCancel(context.Background())

	for i := 0; i < 10; i++ {
		go internal.KafkaPub(*hler.GetPubChan(), &redisRepo, i, &wg1)
		wg1.Add(1)
		go internal.KafkaSub(chan1, i, kafkaSubCtx)
		go internal.WSWriterHandler(chan1, hler.GetMap(), msgTable, &wg2)
		wg2.Add(1)
	}

	// middleware sequence
	r.Handle("/ws", internal.UpgradeValidation(http.HandlerFunc(hler.HandleConn)))

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
	server.Shutdown(ctx) // shutting down ws reader
	kafkaSubCancel()     // shutting down kafka sub
	wg1.Wait()
	wg2.Wait()
	defer cancel()
	// TODO: send shutdown signals to all goroutines

}
