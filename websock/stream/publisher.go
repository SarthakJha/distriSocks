package stream

import (
	"context"
	"encoding/json"
	"time"

	"github.com/SarthakJha/distr-websock/internal/models"
	"github.com/segmentio/kafka-go"
)

func PublishMessage(brokerUrls []string, topic string, partition int, payload models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	w := getKafkaPublisher(brokerUrls, topic)
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = w.WriteMessages(ctx, kafka.Message{
		Partition: partition,
		Key:       []byte("data"),
		Value:     bytePayload,
	})
	if err != nil {
		return err
	}

	return nil
}
