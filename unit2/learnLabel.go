package main

import (
	"fmt"
)

func runLabel() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
		fmt.Println(i)
	}
}

func runGotoLabel1() {
	var i int = 0
loop:
	for i < 10 {
		if i == 4 {
			i++
			goto loop
		}
		fmt.Println(i)
		i++
	}
}

func runGotoLabel2() {
	for i := 0; i < 5; i++ {
		if i == 4 {
			goto loop
		}
		fmt.Println(i)
	}
loop:
	fmt.Println("hellow world")
}
func runControlLabel() {
	//runLabel()
	runGotoLabel1()
	//runGotoLabel2()
}
