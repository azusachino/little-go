package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	client, err := sarama.NewClient([]string{"localhost:9192", "localhost:9292", "localhost:9392"}, config)
	defer client.Close()
	if err != nil {
		panic(err)
	}
	consumer, err := sarama.NewConsumerFromClient(client)

	defer consumer.Close()
	if err != nil {
		panic(err)
	}
	// get partitionId list
	partitions, err := consumer.Partitions("my_topic")
	if err != nil {
		panic(err)
	}

	for _, partitionId := range partitions {
		// create partitionConsumer for every partitionId
		partitionConsumer, err := consumer.ConsumePartition("my_topic", partitionId, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		go func(pc *sarama.PartitionConsumer) {
			defer (*pc).Close()
			// block
			for message := range (*pc).Messages() {
				value := string(message.Value)
				log.Printf("Partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, value)
			}

		}(&partitionConsumer)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:

	}
}
