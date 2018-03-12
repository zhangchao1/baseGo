package baseFunction

import (
	"fmt"
	"log"
	"runtime"
)

func canChangeParameter(a ...string) {
	fmt.Println(a)
}

func canChangeParameterInterface(a ...interface{}) {
	fmt.Println(a)
}

func runLambaFun() {
	lam := func(i string) { fmt.Println(i) }
	lam("hellow word")
}

/*
先执行return以后ret的数值是1，ret++的返回值会是2
*/
func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func add(a int) func(b int) int {
	return func(b int) int {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
		return b + a
	}
}

func runFunc() {
	canChangeParameter("1", "2", "3")
	canChangeParameterInterface("h", "l")
	runLambaFun()
	fmt.Println(f())
	rlam := add(20)
	fmt.Println(rlam(3))
	fmt.Println(rlam(10))
}
