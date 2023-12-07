package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Utwórz bezpośredni exchange
	err = ch.ExchangeDeclare(
		"my_direct_exchange", // nazwa exchange
		"direct",             // typ exchange
		false,                // durable
		false,                // auto-delete
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Utwórz i powiąż kolejki z exchange
	for i := 1; i <= 3; i++ {
		q, err := ch.QueueDeclare(
			fmt.Sprintf("my_queue_%d", i), // nazwa kolejki
			false,                         // durable
			false,                         // delete when unused
			false,                         // exclusive
			false,                         // no-wait
			nil,                           // arguments
		)
		if err != nil {
			log.Fatal(err)
		}

		// Powiąż kolejki z direct exchange używając jednego klucza routingu dla wszystkich
		err = ch.QueueBind(
			q.Name,                           // nazwa kolejki
			fmt.Sprintf("routing_key_%d", i), // klucz routingu
			"my_direct_exchange",             // nazwa direct exchange
			false,                            // no-wait
			nil,                              // arguments
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Producent wysyła wiadomość do direct exchange z określonym kluczem routingu
	err = ch.Publish(
		"my_direct_exchange", // nazwa direct exchange
		"routing_key_1",      // klucz routingu
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Wiadomość do rozgłaszania"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
	}
}
