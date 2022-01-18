package internal

import (
	"log"
	"os"

	"github.com/SarthakJha/distr-websock/models"
	"github.com/SarthakJha/distr-websock/repository"
	"github.com/SarthakJha/distr-websock/stream"
)

func KafkaPub(recvChan chan models.Message, redis *repository.ConnectionRepository, partition int) {
	for {
		select {
		case msg := <-recvChan:
			// query redis for user_id to get pod_id
			val, err := redis.GetWSConnections(msg.RecieverID)
			if err != nil || len(val) == 0 {
				continue
			}
			brokers := []string{
				os.Getenv("KAFKA_BROKER_1"),
				os.Getenv("KAFKA_BROKER_2"),
				os.Getenv("KAFKA_BROKER_3"),
			}
			// publish to kafka to topic pod_id
			err = stream.PublishMessage(brokers, os.Getenv("POD_ID"), partition, msg)
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
	}
}
