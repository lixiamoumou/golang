package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

//执行命令
//先go build
//./66command-line-flags -wat
//./66command-line-flags -h
//./66command-line-flags -word=opt a1 a2 a3 -numb=7
//./66command-line-flags -word=opt a1 a2 a3
//./66command-line-flags -word=opt
//./66command-line-flags -word=opt -numb=7 -fork -svar=flag
