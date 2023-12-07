package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	//"github.com/streadway/amqp"
)

var ch *amqp.Channel

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()
	for {
		// Przykład pobierania listy kolejek w języku Go
		queues, err := ch.QueueDeclarePassive() // lub inna funkcja do pobierania informacji o kolejkach

	}

}
