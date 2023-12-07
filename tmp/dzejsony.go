package main

import (
	"encoding/json"
	"fmt"
)

type MessageMQ struct {
	User string `json:"user"` // transmitter
	Data string `json:"data"` // payload
}

func main() {
	busQueue := "bus"

	busQueue = busQueue

	// Example JSON string
	jsonString := `{"user": "John", "data": "Hello, World!"}`

	// Parse JSON string into a MessageMQ struct
	var msg MessageMQ
	err := json.Unmarshal([]byte(jsonString), &msg)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Access fields of the MessageMQ struct
	fmt.Println("User:", msg.User)
	fmt.Println("Data:", msg.Data)
}
