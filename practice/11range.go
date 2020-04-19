//package main
//
//import (
//	"fmt"
//)
//
//func main(){
//	s := []int{1, 2, 3, 4}
//	sums := 0
//	for _, ss := range s{
//		sums += ss
//	}
//	fmt.Println("sums = ", sums)
//
//	for i, ss := range s{
//		if ss == 3{
//			fmt.Println("index: ", i)
//		}
//	}
//
//	m := map[string]int{"one": 1, "two": 2, "three": 3}
//	for k, v := range m{
//		fmt.Printf("%s->%v\n", k, v)
//	}
//
//	for k, _ := range m{
//		fmt.Println("key: ", k)
//	}
//
//	for i, c := range "go"{
//		fmt.Println(i, c)
//	}
//}

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("s, length: %v, cap: %v\n", len(s), cap(s))

	t := make([]int, len(s), (cap(s)+1)*2)
	copy(s, t)
	fmt.Printf("t, length: %v, cap: %v\n", len(t), cap(t))
	fmt.Printf("s, length: %v, cap: %v\n", len(s), cap(s))

}
