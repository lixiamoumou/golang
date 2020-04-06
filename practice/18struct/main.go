package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Susan"})

	fmt.Println(person{age: 5})

	fmt.Println(&person{"ptr", 50})

	s := person{"Sean", 45}
	fmt.Println(s.name, s.age)

	sp := &s
	fmt.Println(sp.name, sp.age)
	sp.age = 100
	fmt.Println(sp.name, sp.age)

}
