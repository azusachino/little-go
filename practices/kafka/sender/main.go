package main

import (
	"log"

	"github.com/IBM/sarama"
)

var config = sarama.NewConfig()

var localKafka = []string{"127.0.0.1:9093"}

func init() {
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
}

func main() {
	send()
}

func send() {

	msg := sarama.ProducerMessage{}
	msg.Topic = "my-topic"
	msg.Value = sarama.StringEncoder("this is a test")

	client, err := sarama.NewAsyncProducer(localKafka, config)
	if err != nil {
		log.Fatalf("connect failed")
	}
	defer client.Close()

	client.Input() <- &msg

}
