package main

import (
	"fmt"
	"time"
)

func main() {

	////每秒发送一个requests
	//requests := make(chan int, 5)
	//for i := 1; i <= 5; i++ {
	//	requests <- i
	//}
	//close(requests)
	//
	////等价于
	ticker := time.NewTicker(time.Second)
	limiter := ticker.C
	////limiter 通道上持续1s后再发送当前时间
	////limiter := time.Tick(time.Second)
	//
	//for req := range requests {
	//	nowtime := <-limiter
	//	fmt.Println("request", req, nowtime)
	//}

	//前三个请求及时处理
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//如果不起协程，直接for循环，会导致阻塞，一直等待
	go func() {
		for t := range limiter {
			fmt.Println(time.Now())
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
