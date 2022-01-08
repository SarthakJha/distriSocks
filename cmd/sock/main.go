package main

import (
	"github.com/SarthakJha/distr-websock/repository"
)

func main() {
	db := repository.MessageRepository{}
	db.InitMessageConnection()
}
