package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("C:\\Users\\zhangyanhong\\go\\src\\golang\\defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(path string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(file *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("closing")
	err := file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
