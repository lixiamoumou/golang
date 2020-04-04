package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("empty: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("length: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("value: ", b)

	var twoD [3][2]int
	//var twoD = [3][2]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("two dem: ", twoD)

}
