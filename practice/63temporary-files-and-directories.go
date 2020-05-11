package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dname, err := ioutil.TempDir("C:\\Users\\zhangyanhong\\go", "sampledir")
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)

	f, err := ioutil.TempFile("C:\\Users\\zhangyanhong\\go", "sample*.txt")
	check(err)

	fmt.Println("Temp file name:", f.Name())

	// 删除不了？？？
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
}
