package baseGoroutine

import (
	"fmt"
	"time"
)

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-1)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Println(r)
			time.Sleep(delay)
		}
	}
}

func rung1() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Println(n, fibN)
}
