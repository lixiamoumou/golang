package main

import (
	"fmt"
)

func main() {
	//有缓冲通道，允许在没有接受者的情况下，进行缓存capacity数量的数据
	messages := make(chan string, 2)

	messages <- "hello"
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
