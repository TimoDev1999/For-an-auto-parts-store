package messaging

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var conn *amqp091.Connection
var channel *amqp091.Channel

func InitRabbitMQ(url string) {
	var err error
	conn, err = amqp091.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	_, err = channel.QueueDeclare(
		"parts", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}
}

func GetChannel() *amqp091.Channel {
	return channel
}

func CloseRabbitMQ() {
	if err := channel.Close(); err != nil {
		log.Fatalf("Failed to close RabbitMQ channel: %s", err)
	}
	if err := conn.Close(); err != nil {
		log.Fatalf("Failed to close RabbitMQ connection: %s", err)
	}
}
