package stream

import "github.com/segmentio/kafka-go"

func GetKafkaConsumer(brokers []string, consumerGroup string, topic string, partition int64) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		GroupID:   consumerGroup,
		Partition: int(partition),
		Topic:     topic,
	})

	return reader
}

func getKafkaPublisher(brokerUrls []string, topic string) *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokerUrls,
		Topic:   topic,
	})
	return writer
}
