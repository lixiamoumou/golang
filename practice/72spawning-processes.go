package main

import (
	"fmt"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls")
	fmt.Println(string(dateOut))
}
