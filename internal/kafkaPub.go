package internal

import "github.com/SarthakJha/distr-websock/models"

func KafkaPub(recvChan chan models.Message) {
	for {
		select {
		case msg := <-recvChan:
			// query redis for user_id to get pod_id
			// publish to kafka to topic pod_id
		}
	}
}
