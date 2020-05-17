package main

import (
	"fmt"
	"os"
)

func main() {

	//切片中的第一个参数是该程序的路径， 而 os.Args[1:]保存了程序全部的参数
	argWithProg := os.Args
	argWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argWithProg)
	fmt.Println(argWithoutProg)
	fmt.Println(arg)
}

//先go build
//$ go build command-line-arguments.go
//$ ./command-line-arguments a b c d
