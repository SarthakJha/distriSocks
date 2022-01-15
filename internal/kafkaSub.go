package internal

import "github.com/SarthakJha/distr-websock/models"

func KafkaSub(recvChan chan models.Message, writerChan chan models.Message) {
	for {
		select {
		case recvMsg := <-recvChan:
			// put into writer chan
			// set status to SENT

		}
	}
}
