package main

import (
	"fmt"
	"time"
)

func main() {
	start1 := time.Now()
	message := make(chan string)
	signal := make(chan bool)

	go func() {
		//time.Sleep(2 *time.Second)
		message <- "func hello"
		//signal <- true
	}()

	go func() {
		//time.Sleep(2 *time.Second)
		signal <- true
	}()

	select {
	case msg1 := <-message:
		fmt.Println("received: ", msg1)
	default:
		fmt.Println("no received message.", time.Since(start1))
	}

	start2 := time.Now()
	msg := "hello"
	select {
	case message <- msg:
		fmt.Println("send: ", msg)
	default:
		fmt.Println("no send message.", time.Since(start2))
	}

	select {
	case msg2 := <-message:
		fmt.Println("received message: ", msg2, time.Since(start1))
	case sig := <-signal:
		fmt.Println("received signal: ", sig, time.Since(start1))
	default:
		fmt.Println("received noting.", time.Since(start1))
	}
}
