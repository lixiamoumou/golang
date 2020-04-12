//package main
//
//import (
//	"fmt"
//)
//
//func main(){
//	jobs := make(chan int)
//	done := make(chan bool)
//
//	go func() {
//		for {
//			j, more := <- jobs
//			if more{
//				fmt.Println("received job", j)
//			}else{
//				fmt.Println("received all jobs")
//				done <- true
//				return
//			}
//		}
//	}()
//
//	for i := 1; i < 4; i++{
//		jobs <- i
//		fmt.Println("send job ", i)
//	}
//	close(jobs)
//	fmt.Println("send all jobs")
//
//	<-done
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//func f1(in chan int) {
//	fmt.Println(<-in)
//}
//
//func main() {
//	out := make(chan int)
//	//out <- 2在goroutine之前会报错，deadlock
//	//out <- 2
//	go f1(out)
//	out <- 2
//}

package main

import "fmt"

func main() {
	//报错deadlock
	// chan绝对不是全局变量，一个全局变量，可以在同一函数体内重复读写，但对无缓冲chan而言是不可以，
	// 原因在同一goroutine内对同一chan读写时，存在读或写阻塞面临切换上下文，另一个对应的永远没执行机会，
	//ch := make(chan int)
	ch := make(chan int, 1)
	ch <- 5
	fmt.Println(<-ch)
}
