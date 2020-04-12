package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 3")
	}

	fmt.Println(time.Since(start))
}
