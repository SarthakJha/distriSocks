package internal

import "github.com/SarthakJha/distr-websock/models"

func KafkaPub(recvChan chan models.Message) {
	for {
		select {
		case msg := <-recvChan:
			//
		}
	}
}
