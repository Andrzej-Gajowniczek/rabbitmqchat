package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var licznik int

func main() {
	// Tworzymy context z funkcją anulowania (cancel function)
	ctx, cancel := context.WithCancel(context.Background())

	// WaitGroup, aby poczekać na zakończenie goroutine
	var wg sync.WaitGroup

	// Uruchamiamy goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				// Wyjście z goroutine, gdy context zostanie anulowany
				fmt.Println("Goroutine została zakończona")
				return
			default:
				// Wykonuj jakieś operacje
				fmt.Println("Goroutine wykonuje operacje", licznik)

				// Sprawdzamy warunek, po którym chcemy zakończyć goroutine
				if licznik > 9 {
					cancel() // Anuluj context, co spowoduje zakończenie goroutine
					return
				}

				// Odczekaj przed kolejnym wykonaniem operacji
				time.Sleep(time.Second)
				licznik++
			}
		}
	}()

	// Czekamy na zakończenie goroutine
	wg.Wait()
	fmt.Println("Program zakończony")
}

// warunek to warunek, po którym chcemy zakończyć goroutine
//var warunek bool
