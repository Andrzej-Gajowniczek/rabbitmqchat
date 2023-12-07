package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Wiad struct {
	FROM string `json: from`
	TO   string `json: to`
	INFO string `json: data`
}

func main() {
	var info Wiad
	pobieracz := bufio.NewScanner(os.Stdin)
jeszczeRaz:
	pobieracz.Scan()

	eureka := pobieracz.Text()
	log.Println(eureka)
	if eureka == "" {
		goto jeszczeRaz
	}
	words := strings.Fields(eureka)
	for i, v := range words {
		switch i {
		case 0:
			info.FROM = v
		case 1:
			info.TO = v
		case 2:
			info.INFO = strings.Join(words[2:], "_")
		}
	}
	fmt.Println(info)
	dzejzon, err := json.Marshal(info)
	if err == nil {
		fmt.Println(string(dzejzon))
		fmt.Printf("to jest dzejzon:%s", dzejzon)
	}
}
