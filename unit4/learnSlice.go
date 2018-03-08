package unit4

import (
	"fmt"
)

func basicSlice() {
	var x = []int{2, 3, 5, 7, 11, 12, 14, 14}
	y := x[3:4]
	fmt.Println(len(x), cap(x), len(y), cap(y))

	y2 := y[0:5]
	fmt.Println(len(y2), cap(y2))
}

func diffNewAndMake() {
	var b *[]int = new([]int)
	*b = []int{1, 2, 3}
	fmt.Println(b)
	c := b
	*c = []int{6, 7, 89}
	fmt.Println(b, c)

	var b1 = make([]string, 10, 50)
	fmt.Println(b1)
	a1 := b1
	a1[6] = "m"
	fmt.Println(a1, b1)
}

func runForRange() {
	var a = [...]int{10, 20, 30}
	for i, _ := range a {
		a[i] *= 2
	}
	fmt.Println(a)
}

func runAppend() {
	var b = []int{1, 2, 4, 5}
	fmt.Println(len(b), cap(b), b)
	c := append(b, 2, 4, 5, 6, 7)
	fmt.Println(len(c), cap(c), c)

	var e = []int{1, 2, 3, 4}
	for i := 0; i < 10; i++ {
		fmt.Println(len(e), cap(e), e)
		e = append(e, i)

	}
}

func runCopy() {
	var a = []int{4, 5, 6, 7, 8, 9, 10}
	fmt.Println(len(a), cap(a))
	var b = []int{2, 3, 4, 5}
	copy(a, b)
	fmt.Println(len(a), cap(a), a, b)
}

func runSlice() {
	//basicSlice()
	//diffNewAndMake()
	//runForRange()
	//runAppend()
	runCopy()
}
