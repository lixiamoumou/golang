package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dirname := "subdir"
	//err为nil，则有文件
	if _, err := os.Stat(dirname); err == nil {
		os.RemoveAll(dirname)
	}
	// 在当前工作目录Go下创建文件夹
	err := os.Mkdir(dirname, 0755)
	check(err)

	defer os.RemoveAll(dirname)

	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir\\file1.txt")

	err = os.MkdirAll("subdir\\parent\\child", 0755)
	check(err)

	createEmptyFile("subdir\\parent\\file2.txt")
	createEmptyFile("subdir\\parent\\file3.txt")
	createEmptyFile("subdir\\parent\\child\\file4.txt")

	r, err := ioutil.ReadDir("subdir\\parent")
	check(err)
	fmt.Println("List subdir/parent")
	//返回的第一个参数是下标
	for _, entry := range r {
		//fmt.Println(a)
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir\\parent\\child")
	check(err)

	r, err = ioutil.ReadDir(".")
	check(err)
	fmt.Println("List subdir/parent/child")
	//返回的第一个参数是下标
	for _, entry := range r {
		//fmt.Println(a)
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)
	r, err = ioutil.ReadDir(".")
	fmt.Println("List Go")
	for _, entry := range r {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// walk方法会遍历路径下的所有文件(包含路径)并对每一个目录和文件都调用walkFunc方法。
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
