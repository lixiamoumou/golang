package main

import (
	"fmt"
	"time"
)

func main() {
	tickers := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-tickers.C:
				fmt.Println("Tick at: ", t)
			}
		}
	}()

	time.Sleep(20 * time.Second)
	tickers.Stop()
	done <- true

	fmt.Println("Ticker stopped")

}
