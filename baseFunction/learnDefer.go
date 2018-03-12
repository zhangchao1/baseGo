package baseFunction

import (
	"fmt"
)

func runDefer() {
	fmt.Println("start")
	defer sayDefer("end")
	fmt.Println("next")
}

func sayDefer(a string) {
	fmt.Println(a)
}

func runDeferOrder() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func runDefer() {
	//runDefer()
	//fmt.Println(1)
	runDeferOrder()
}
