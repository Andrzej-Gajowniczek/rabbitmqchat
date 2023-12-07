package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Adres RabbitMQ
	rabbitMQAddr := "amqp://guest:guest@localhost:5672/"

	// Utwórz połączenie z RabbitMQ
	conn, err := amqp.Dial(rabbitMQAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Utwórz kanał
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Pobierz listę aktywnych kolejek
	queueList, err := ch.()
	if err != nil {
		log.Fatal(err)
	}

	// Wydrukuj listę aktywnych kolejek
	fmt.Println("Aktywne kolejki:")
	for _, q := range queueList.Queues {
		fmt.Println(q.Name)
	}
}
