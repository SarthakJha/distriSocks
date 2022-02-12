package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/SarthakJha/distr-websock/internal/models"
	"github.com/SarthakJha/distr-websock/internal/utils"
	"github.com/SarthakJha/distr-websock/stream"
)

func KafkaSub(writerChan chan models.Message, consumingPartition int, ctx context.Context, conf utils.Config) {
	brokers := utils.ResolveHeadlessServiceDNS(conf.KAFKA_SERVICE, "kafka")
	reader := stream.GetKafkaConsumer(brokers, "something", os.Getenv("POD_ID"), int64(consumingPartition))
	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		// listen for elements then write to writer chan
		select {
		// routine shutdown case
		case <-ctx.Done():
			if err := reader.Close(); err != nil {
				fmt.Println(err.Error())
			}
			close(writerChan)
			cancel()
			return

		default:
			msg, err := reader.ReadMessage(ctx1)
			if err != nil {
				continue
			}
			var message models.Message
			err = json.Unmarshal(msg.Value, &message)
			if err != nil {
				continue
			}
			writerChan <- message

		}

	}
}
