package main

import (
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()
	//Timer结构里面的C是单向管道，只能读不能写
	//NewTimer()创建一个新的计时器，该计时器将在其通道上至少持续指定时长之后发送当前时间。
	timer1 := time.NewTimer(2 * time.Second)

	//NewTimer()往C通道中写Now(),下一行代码从通道中读取数据
	now1 := <-timer1.C
	duration1 := time.Since(starttime)
	fmt.Println(now1)
	fmt.Println("timer1 is over, ", duration1)

	starttime2 := time.Now()
	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer2.C
		duration2 := time.Since(starttime2)
		fmt.Println("timer2 is over, ", duration2)
	}()
	duration3 := time.Since(starttime2)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 is stop, ", duration3)
	}
}
