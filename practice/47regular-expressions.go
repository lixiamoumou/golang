package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	//func MatchString(pattern string, s string) (matched bool, err error)
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//Compile解析正则表达式，如果成功，则返回可用于与文本匹配的Regexp对象。
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.MatchString("peach"))

	//FindString returns a string holding the text of the leftmost match in s of the regular expression.
	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.FindAllStringSubmatch("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	//字符串强转为byte切片
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
