package main

import (
	"log"
	"net/http"

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

	http.HandleFunc("/ws", hler.HandleConn)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
