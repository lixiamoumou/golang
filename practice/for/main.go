package main

import (
	"fmt"
)

func main(){
	// first kind of for loop
	i := 1
	for i<=3{
		fmt.Println(i)
		i++
	}

	//second kind of for loop
	for j:=7;  j<=9; j++{
		fmt.Println(j)
	}

	//third kind of for loop
	for{
		fmt.Println("loop")
		break
	}

	//continue
	for n:=0; n<=5; n++{
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	}
}
