package main

import (
	"fmt"
)

func runInt() {
	var a int32
	var b int64
	a = 16
	// b = a + a 错误的用法，需要进行转换，才可以进行
	b = int64(a)
	//要进行转换
	fmt.Println(a, b)
}

func main() {
	runInt()
}
