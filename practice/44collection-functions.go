package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsf := make([]string, len(vs))
	for i, v := range vs {
		vsf[i] = f(v)
	}
	return vsf
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	//输出指定字符串的下标
	fmt.Println(Index(strs, "pear"))
	//输出是否包含指定字符串
	fmt.Println(Include(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	//输出是否有满足函数的字符串
	fmt.Println(Any(strs, func(i string) bool {
		return strings.HasPrefix(i, "p")
	}))
	//输出是否都满足函数
	fmt.Println(All(strs, func(i string) bool {
		return strings.HasPrefix(i, "p")
	}))
	//对切片进行过滤
	fmt.Println(Filter(strs, func(i string) bool {
		return strings.Contains(i, "e")
	}))
	//使用函数对切片进行映射
	fmt.Println(Map(strs, strings.ToUpper))
}
