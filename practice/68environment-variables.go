package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		//n表示分割成几部分
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair)
		fmt.Println(pair[0])
	}
}
