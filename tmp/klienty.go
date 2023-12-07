package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func check4errors(s string, err error) {

	if err != nil {
		log.Println(s, err)
	}
}
func main() {

	// flagi
	var hostname string
	if len(os.Args) > 1 {
		//flags := os.Args[1:]

		flag.StringVar(&hostname, "host", "localhost", "--host hostname")
		flag.Parse()
	}
	//Połączenie do RabbitMQ
	uri := "amqp://guest:guest@" + hostname + ":5672/"
	conn, err := amqp.Dial(uri)
	check4errors("Failed to connect to RabbitMQ: %v", err)
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	check4errors("Failed to open a channel:", err)
	defer ch.Close()

	//zapytanie użytkownika o nazwę
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Podaj imie albo nick: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Printf("username:%s\n", username)

	//deklaracja kolejki ekskluzywnej dla kanału username:
	q, err := ch.QueueDeclare(
		username, // unikalna nazwa kolejki dla użytkownika
		false,    // durable
		false,    // delete when unused
		true,     // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	//deklaracja fanout
	err = ch.ExchangeDeclare(
		"broadcast", // nazwa exchange
		"fanout",    // typ exchange
		false,       // durable
		false,       // auto-delete
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	check4errors("fanout", err)

	//bindowanie eksklusiwa do fanout
	err = ch.QueueBind(
		q.Name,      // nazwa kolejki
		"",          // routing key (puste dla fanout)
		"broadcast", // nazwa fanout exchange
		false,       // no-wait
		nil,         // arguments
	)

	go func() {
		for {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			message := scanner.Text()
			err = ch.Publish(
				"broadcast", // nazwa fanout exchange
				"",          // routing key (puste dla fanout)
				false,       // mandatory
				false,       // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(message),
				},
			)
			check4errors("Publikacja broadcast", err)
		}
	}()

	// konsumowanie
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	check4errors("odbieranie msg", err)

	// Handle incoming messages
	go func() {
		for msg := range msgs {
			fmt.Printf("my username is%s :%s\n", q.Name, msg.Body)
		}
	}()

	for {
	}
}
