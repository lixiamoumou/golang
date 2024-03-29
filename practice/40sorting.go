package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{7, 2, 4}
	if !sort.IntsAreSorted(ints) {
		sort.Ints(ints)
	}
	fmt.Println("Ints: ", ints)

}
