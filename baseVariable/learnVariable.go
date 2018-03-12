package baseVariable

import (
	"fmt"
	"os"
)

var (
	v1 int
	v2 bool
	v3 string
	v4 float64
)

var v5 = 123

//并行赋值
var v9, v10, v11 = 123, "1", true

/*当一个变量被声明之后，
系统自动赋予它该类型的零值：
int 为 0，float 为 0.0，bool 为 false，string 为空字符串，
指针为 nil。
记住，所有的内存在 Go 中都是经过初始化的。
*/

//包内部的变量是大写的，可以外部调用的

//全局变量允许声明，但是不使用

//交换变量的数值 a,b=b,a不用使用第三方变量

func changeV5() {
	v5 = 234
}

func runVariable() {
	v7 := v5
	v8 := &v5
	fmt.Println(v1, v2, v4)
	fmt.Sprint(v3)
	fmt.Println(os.Getenv("PATH"))
	//获取系统的环境变量
	fmt.Println(v7, v8)

	//全局变量的改变
	fmt.Println(v5)
	changeV5()
	fmt.Println(v5)
}
