package main

import (
	"fmt"
	"os"
)

func main() {

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
