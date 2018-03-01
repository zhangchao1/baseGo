// package unit1

// import (
// 	"fmt"
// )

// type ByteSize float64

// //iota

// //参考文档地址 https://golang.org/ref/spec#Iota
// const (
// 	KB ByteSize = 1 << iota
// 	MB
// 	GB
// )

// //默认复制是0
// const (
// 	NZ = iota
// 	NB
// 	NM
// )

// //FRI定义的时候，会根据上面的常量来进行判断
// const (
// 	WE = 1
// 	FRI
// 	SES
// )

// //SES的输出是3，

// func runConstant() {
// 	fmt.Println(KB, MB, GB)
// 	fmt.Println(WE, FRI, SES)
// 	fmt.Println(NZ, NB, NM)
// }
