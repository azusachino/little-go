package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		log.Fatalf("fail to connect kafka: %v", err)
	}
	delCh := make(chan kafka.Event, 10000)
	defer close(delCh)
	topic := "default-topic"
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte("Test kafka by golang"),
	}, delCh)

	chOut := <-delCh
	messageReport := chOut.(*kafka.Message)
	if messageReport.TopicPartition.Error != nil {
		log.Fatalf("err occured")
	} else {
		log.Println("success")
	}
}
