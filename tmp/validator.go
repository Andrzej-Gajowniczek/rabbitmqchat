package main

import (
	"fmt"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

type MessageMQ struct {
	User string `json:"user" validate:"required"`
	Data string `json:"data" validate:"required"`
	Year int    `json:"year"`
}

func main() {
	// Przykładowy JSON do walidacji
	jsonData := `{"user": "John", "data": "Hello, World!"}`

	// Utwórz schemat JSON na podstawie struktury MessageMQ
	schemaLoader := gojsonschema.NewReferenceLoader("file://path/to/your/schema.json")

	// Utwórz dane JSON do walidacji
	documentLoader := gojsonschema.NewStringLoader(jsonData)

	// Waliduj dane JSON względem schematu
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		log.Fatal(err)
	}

	// Sprawdź, czy walidacja zakończyła się sukcesem
	if result.Valid() {
		fmt.Println("JSON jest poprawny!")
	} else {
		fmt.Println("JSON jest niepoprawny.")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
