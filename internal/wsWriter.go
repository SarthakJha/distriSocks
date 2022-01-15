package internal

import (
	"log"

	"github.com/SarthakJha/distr-websock/models"
)

func WSWriterHandler(writeBuff chan models.WSWriterBuffer) {
	for {
		select {
		case msg := <-writeBuff:
			// writing back the message to the socket
			err := msg.WSConnection.WriteJSON(msg.Payload)
			if err != nil {
				// if socket doesnt exist or inactive then logging
				log.Println("error: ", err.Error())
				msg.WSConnection.Close()

				// TODO: remove this connection key from redis
				// TODO: mark message as sent
			}
		}
	}
}
