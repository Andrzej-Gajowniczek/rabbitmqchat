package main

import (
	"flag"
	"fmt"
	"os"
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
}
