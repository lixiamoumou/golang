package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	//要产生不同的数字序列，需要给定一个不同的种子。
	//使用时间作为种子，每次运行代码生成的数值不同
	fmt.Println(time.Now().UnixNano())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	r2 := rand.New(rand.NewSource(42))
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	r3 := rand.New(rand.NewSource(42))
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
