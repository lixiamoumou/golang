package main

import (
	"fmt"
)

func sum(num ...int) {
	fmt.Print(num, " ")

	sums := 0
	for _, i := range num {
		sums += i
	}

	fmt.Println(sums)
}

func main() {
	sum(1, 2, 3)
	sum(9, 8, 7, 6)

	num := []int{1, 2, 3, 4}
	sum(num...)

}
