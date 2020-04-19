package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)

	queue <- "One"
	queue <- "Two"
	//没有close的话，会报错fatal error: all goroutines are asleep - deadlock!
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
