package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	//set hostname of rabbitMQ service to connect to
	var hostname string = "127.0.0.1"
	var port string = "5672"
	if len(os.Args) > 1 {
		flag.StringVar(&hostname, "host", "localhost", "--host hostname")
		flag.StringVar(&port, "port", "5672", "--port 5672")
		flag.Parse()
	}
	uri := "amqp://guest:guest@" + hostname + ":" + port + "/"
	fmt.Println(uri)

	//try to connect
	conn, err := amqp.Dial(uri)
	check4errors("Connecting RabbitMQ:", err)
	defer conn.Close()
}

func check4errors(s string, err error) {
	if err != nil {
		log.Println(s, "\x1b[1;31mError:", err)
		os.Exit(-1)
	} else {
		log.Println(s, "\x1b[1;92m"+"Successful")
	}
}
