package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MessageMQ struct {
	User string `json:"user"` // transmitter
	Data string `json:"data"` // payload
	Year int    `json:"year"`
}

func main() {

	struna := os.Args[1]
	// Create an instance of the struct
	message := MessageMQ{
		User: "John",
		Data: "Hello, World!",
		Year: 2001,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Print the JSON data as a string
	fmt.Println(string(jsonData))
	var nieserlane map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &nieserlane)
	if err != nil {
		fmt.Println("jsonData nawala")
	}
	fmt.Println(nieserlane)
	var jarek map[string]interface{}
	err = json.Unmarshal([]byte(struna), &jarek)
	if err != nil {
		fmt.Println("struna nawala", jarek, struna)

	}
	fmt.Println(jarek)
}
