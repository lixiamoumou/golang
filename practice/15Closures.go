package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt01 := intSeq()
	fmt.Println(nextInt01())

	fmt.Println(intSeq()())
	fmt.Println(intSeq()())
}
