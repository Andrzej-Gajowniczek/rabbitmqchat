package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Przykładowy JSON
	jsonData := `{"user": "John", "data": {"age": 30, "city": "New York"}}`

	// Utwórz mapę, aby przechować dynamiczny JSON
	var data map[string]interface{}

	// Parsuj JSON do mapy
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Błąd parsowania JSON:", err)
		return
	}

	// Wyświetl dane z mapy
	fmt.Println("User:", data["user"])

	// Odwołanie do zagnieżdżonych danych
	if nestedData, ok := data["data"].(map[string]interface{}); ok {
		fmt.Println("Age:", nestedData["age"])
		fmt.Println("City:", nestedData["city"])
	}
}
