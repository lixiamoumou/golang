package main

import (
	"fmt"
)

func plus(a, b int) int {
	return a + b
}

func plusThree(a, b, c int) int {
	return a + b + c
}

func main() {
	a, b, c := 1, 2, 3

	addition := plus(a, b)
	fmt.Println("a + b = ", addition)
	additionthree := plusThree(a, b, c)
	fmt.Println("a + b + c = ", additionthree)
}
