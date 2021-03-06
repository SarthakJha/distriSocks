package internal

import (
	"log"
	"os"
	"sync"

	"github.com/SarthakJha/distr-websock/internal/models"
	"github.com/SarthakJha/distr-websock/internal/utils"
	"github.com/SarthakJha/distr-websock/repository"
	"github.com/SarthakJha/distr-websock/stream"
)

func KafkaPub(recvChan chan models.Message, redis *repository.ConnectionRepository, partition int, wg *sync.WaitGroup, conf utils.Config) {
	defer wg.Done() // graceful shutdown
	// this will automatically shut when producing side of channel shuts
	for msg := range recvChan {
		// query redis for user_id to get pod_id
		val, err := redis.GetWSConnections(msg.RecieverID)
		if err != nil || len(val) == 0 {
			continue
		}
		brokers := utils.ResolveHeadlessServiceDNS(conf.KAFKA_SERVICE, "kafka") // publish to kafka to topic pod_id
		err = stream.PublishMessage(brokers, os.Getenv("POD_ID"), partition, msg)
		if err != nil {
			log.Fatalln(err.Error())
		}

	}
}
