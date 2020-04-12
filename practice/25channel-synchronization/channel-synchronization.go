package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second)
	fmt.Println("done")

	//通过通道done往另一个协程(main)发送信息true
	done <- true
}

func main() {
	//创建一个通道
	done := make(chan bool)
	//开启一个协程
	go worker(done)

	//接收worker协程中的信息true
	<-done
}
