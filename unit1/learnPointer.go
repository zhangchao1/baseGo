package unit1

import (
	"fmt"
)

func runPonter() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var ip *int = &i1
	*ip = 7
	fmt.Printf("The value at memory location %p is %d\n", ip, *ip)
	fmt.Println(i1)
}
