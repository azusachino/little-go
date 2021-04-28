package main

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

var localKafka = []string{"127.0.0.1:9093"}

func main() {
	consume()
}

func consume() {
	consumer, err := sarama.NewConsumer(localKafka, nil)
	if err != nil {
		log.Fatalf("error \n")
	}
	partitionList, err := consumer.Partitions("my-topic")
	if err != nil {
		log.Fatalf("err %v", err)
	}
	for p := range partitionList {
		pc, err := consumer.ConsumePartition("my-topic", int32(p), sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("error %v", err)
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				log.Default().Printf("received %v \n", msg.Value)
			}
		}(pc)
	}
	time.Sleep(10 * time.Second)
}
