package main

import (
	"errors"
	"fmt"
)

func f1(num int) (int, error) {
	if num == 24 {
		return -1, errors.New("can't work with 42")
	}
	return num + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(num int) (int, error) {
	if num == 24 {
		return -1, &argError{num, "can't work with it"}
	}
	return num + 3, nil
}

func main() {
	for _, v := range []int{7, 24} {

		if flag, err := f1(v); err != nil {
			fmt.Println("f1 faild, ", err)
		} else {
			fmt.Println("f1 successed, ", flag)
		}
	}

	for _, v := range []int{7, 24} {
		if flag, err := f2(v); err != nil {
			fmt.Println("f2 faild, ", err)
		} else {
			fmt.Println("f2 successed, ", flag)
		}
	}

	//err is argError struct
	_, err := f2(24)
	//ae是err的值，ok表示是否有值
	//只有interface{}类型才能进行断言，也就是使用类似err.(*argError)格式判断err是否为*argError类型
	//builtin.go中定义error,是个接口
	// type error interface {
	//	Error() string
	//}
	if ae, ok := err.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	a := 7
	//转换成interface{}，否则报错，invalid type assertion: a.(int) (non-interface type int on left)
	if num, ok := interface{}(a).(int); ok {
		fmt.Println(num)
	}
}
