package main

import (
	"fmt"
	"time"
)

//jobs通道用来从通道中读取数据，results通道用来往通道中写数据
func pool(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Println("pool ", id, "start job ", i)
		time.Sleep(time.Second)
		fmt.Println("pool ", id, "finish job ", i)
		results <- i
	}
}

func main() {
	const numJobs = 5
	//有缓冲通道，允许在没有接收者的情况下，可以缓存capacity数量的数据
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//开始时间
	starttime := time.Now()

	//起协程
	for p := 1; p <= 3; p++ {
		go pool(p, jobs, results)
	}

	//往jobs通道中写入数据，等价于要处理的jobs的数量
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	//关闭jobs通道
	close(jobs)

	//从results通道中读取结果
	for r := 1; r <= numJobs; r++ {
		fmt.Println(<-results)
	}

	//理论上，如果没有协程的话，完成5个jobs共需要5s
	//到读取到results结果需要的时间
	duration := time.Since(starttime)
	fmt.Println(duration)
}
