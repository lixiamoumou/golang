package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	//秒数
	fmt.Println(secs)
	//毫秒
	fmt.Println(millis)
	//纳秒
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
