package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	//Printf通过os.Stdout输出
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 123)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 33)

	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "h")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	//Sprintf返回字符串
	s := fmt.Sprintf("a %s\n", "string")
	fmt.Printf(s)

	//Fprintf输出到io.Writers,必须实现了 io.Writer 接口
	fmt.Fprintf(os.Stdout, "an %s\n", "error")
}
