package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 利用ioutil.ReadFile直接从文件读取到[]byte中
	data, err := ioutil.ReadFile("C:\\Users\\zhangyanhong\\go\\src\\golang\\defer.txt")
	check(err)
	fmt.Println(string(data))

	// 先从文件读取到file中，在从file读取到buf, buf在追加到最终的[]byte
	f, err := os.Open("C:\\Users\\zhangyanhong\\go\\src\\golang\\defer.txt")
	check(err)
	defer f.Close()

	b1 := make([]byte, 7)
	// n1是文件中实际长度,最大为b1的长度
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	//string := "data\n"
	//fmt.Println(string,len(string))

	//0表示相对于文件的原点，1表示相对于当前偏移量，2表示相对于结尾
	//从offset开始进行读取，读取长度根据b2的长度决定
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n3, o3)
	fmt.Printf("%v\n", string(b3[:n3]))

	_, err = f.Seek(0, 0)
	check(err)

	//先从文件读取到file, 在从file读取到Reader中，从Reader读取到buf, buf最终追加到[]byte
	r4 := bufio.NewReader(f)
	//Peek返回输入流的下n个字节，不会移动读取位置
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	b5, err := r4.Peek(2)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b5))
}
