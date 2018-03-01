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

//复数的运算
func runComplex() {
	var c complex64
	c = 10 + 50i
	fmt.Println(c)
}

//位运算
func runTwoBit() {
	b1 := 1 & 2
	b2 := 1 | 2
	b3 := 1 ^ 2
	b4 := 0 &^ 1
	fmt.Println(b1, b2, b3, b4)
}

func runOneBit() {
	bb1 := ^2
	bb2 := ^-2
	fmt.Println(bb1, bb2)
}

func main() {
	//runInt()
	//runComplex()
	//runTwoBit()
	runOneBit()
}
