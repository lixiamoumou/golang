package main

import (
	"fmt"
)

//定义一个只写通道
//pings是接收数据的通道(只写)
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//定义一个只读和只写通道
//pings接收数据的通道(只读)，pongs从通道接收数据(只写)
func pong(pongs chan<- string, pings <-chan string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message")
	pong(pongs, pings)
	fmt.Println(<-pongs)
}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//var echo chan string
//var receive chan string
//
//// 定义goroutine 1
//func Echo() {
//	time.Sleep(1*time.Second)
//	echo <- "咖啡色的羊驼"
//}
//
//// 定义goroutine 2
//func Receive() {
//	temp := <- echo // 阻塞等待echo的通道的返回
//	receive <- temp
//}
//
//
//func main() {
//	echo = make(chan string)
//	receive = make(chan string)
//
//	go Echo()
//	go Receive()
//
//	getStr := <-receive   // 接收goroutine 2的返回
//
//	fmt.Println(getStr)
//}
