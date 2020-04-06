package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width = 100
	r.height = 100
	return r.width * r.height
}

func (r rect) perim() int {
	r.width = 100
	r.height = 100
	return 2 * (r.height + r.width)
}

func main() {
	//指针隐式解引用
	//调用方法时，Go 会自动处理值和指针之间的转换

	r := rect{10, 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	ptr := &r
	fmt.Println("area1: ", ptr.area())
	fmt.Println("perim1: ", ptr.perim())

}
