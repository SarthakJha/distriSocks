package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SarthakJha/distr-websock/internal"
	"github.com/SarthakJha/distr-websock/repository"
)

func main() {
	msgTable := repository.MessageRepository{}
	usrTable := repository.UserRepository{}
	redisRepo := repository.ConnectionRepository{}
	hler := internal.Chans{}
	hler.InitChan(10)

	redisRepo.InitConnectionRepository()
	msgTable.InitMessageConnection()
	usrTable.InitUserConnection()

	// make 2 channels
	// create go routines

	http.HandleFunc("/ws", hler.HandleConn)

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
