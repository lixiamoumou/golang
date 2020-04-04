package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("empty: ", s)

	s[0] = "one"
	s[1] = "two"
	s[2] = "three"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("length: ", len(s))

	s = append(s, "four")
	s = append(s, "five", "six")

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	sc1 := s[2:5]
	fmt.Println("3-5: ", sc1)

	sc2 := s[:5]
	fmt.Println("1-5", sc2)

	sc3 := s[2:]
	fmt.Println("3-6", sc3)

	t := []string{"a", "b", "c"}
	fmt.Println("slice: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)
}
