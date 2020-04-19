package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height int
}

type circle struct {
	radius int
}

func (r rect) area() float64 {
	return float64(r.height * r.width)
}

func (r rect) perim() float64 {
	return float64(2 * (r.width + r.height))
}

func (c circle) area() float64 {
	return math.Pi * float64(c.radius*c.radius)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * float64(c.radius)
}

func measure(geo geometry) {
	fmt.Println(geo)
	fmt.Println(geo.area())
	fmt.Println(geo.perim())
}

func main() {
	rect := rect{2, 3}
	circle := circle{5}

	fmt.Printf("%T\n", rect)
	fmt.Printf("%T\n", circle)

	measure(rect)
	measure(circle)
}
