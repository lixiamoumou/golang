package main

import (
	"fmt"
)

func zeroval(val int) {
	val = 0
	fmt.Println("zeroval inner address: ", &val)
}

func zeroptr(ptr *int) {
	*ptr = 0
	fmt.Println("zeroptr inner address: ", ptr)
}

func main() {
	num := 7
	fmt.Println("initial:", num)

	zeroval(num)
	fmt.Println("zeroval: ", num)

	zeroptr(&num)
	fmt.Println("zeroptr: ", num)
	fmt.Println("address: ", &num)
}
