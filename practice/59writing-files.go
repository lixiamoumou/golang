package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	//f the file does not exist, WriteFile creates it with permissions perm; otherwise WriteFile truncates it before writing.
	//某位是1时，则把这位的perm属性关闭(disable)
	//某位是0时，则把这位的perm属性打开(enable)
	//   owner   group   other
	//0 - rwx  -  rwx  -  rwx
	//例如
	//$ umask 0222  # 对owner, group,和other的写属性都关闭。
	//$ umask 0124  # 对owner关闭读属性，对group关闭写属性，对other关闭读和写属性
	err := ioutil.WriteFile("C:\\Users\\zhangyanhong\\go\\src\\golang\\test1.txt", d1, 0644)
	check(err)

	f, err := os.Create("C:\\Users\\zhangyanhong\\go\\src\\golang\\test2.txt")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	// 自动翻译ascll码
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	//调用 Sync 将缓冲区的数据写入硬盘
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
