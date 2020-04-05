package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	fmt.Println("empty: ", m)

	m["k1"] = 1
	m["k2"] = 7
	m["k3"] = 17
	fmt.Println("set: ", m)

	k1 := m["k1"]
	fmt.Println("get k1: ", k1)

	length := len(m)
	fmt.Println("length: ", length)

	delete(m, "k2")
	fmt.Println("del k2: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	m2 := map[string]int{"one": 1, "two": 2}
	fmt.Println("m2: ", m2)
}
