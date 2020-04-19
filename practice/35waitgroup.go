package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {

	defer fmt.Println(wg)
	defer wg.Done()

	fmt.Println(wg)
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	//counter： 当前还未执行结束的goroutine计数器
	//waiter count: 等待goroutine-group结束的goroutine数量，即有多少个等候者
	//semaphore: 信号量。gr执行结束为1，否则为0

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
