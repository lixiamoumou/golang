package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	fmt.Printf("%T\n", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	dir, filename := filepath.Split(p)
	fmt.Println("Split(p):", dir, filename)

	fmt.Println(filepath.IsAbs("..\\dir\\file"))
	fmt.Println(filepath.IsAbs("C:\\Users\\zhangyanhong\\go\\src\\golang\\practice"))

	filename = "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	//清除扩展名后的值
	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
