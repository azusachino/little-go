package main

import (
	"context"
	"github.com/azusachino/little-go/practices/rabbitmq"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@" + rabbitmq.IP + ":5672")
	failOnError(err, "failed to connect")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "failed to create channel")
	defer channel.Close()

	withSelect(channel)

}

func withSelect(channel *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go sendCtx(ctx, channel)

	time.Sleep(time.Second * 3)
	cancel() // trigger <- ctx.Done()
}

func withWaitGroup(channel *amqp.Channel) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sendWg(&wg, channel)
	}
	wg.Wait()
}

func sendWg(wg *sync.WaitGroup, ch *amqp.Channel) {

	defer wg.Done()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare queue")

	body := "Hello RabbitMQ " + time.Now().String()
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{ContentType: "text/plain",
			Body: []byte(body)},
	)
	failOnError(err, "failed to publish")

}
func sendCtx(ctx context.Context, ch *amqp.Channel) {
	for {
		select {
		default:
			q, err := ch.QueueDeclare(
				"hello",
				false,
				false,
				false,
				false,
				nil,
			)
			failOnError(err, "failed to declare queue")

			body := "Hello RabbitMQ " + time.Now().String()
			err = ch.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{ContentType: "text/plain",
					Body: []byte(body)},
			)
			failOnError(err, "failed to publish")
		case <-ctx.Done():
			return
		}
	}

}
