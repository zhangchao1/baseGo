package baseSliceAndArray

import (
	"fmt"
)

func valueArray(a [3]int) {
	fmt.Println(a)
	a = [3]int{40, 50, 60}
}

func pointerArray(a *[3]int) {
	fmt.Println(a)
	*a = [3]int{50, 60, 70}
	fmt.Println(&a)
}

func comparePointerAndValue() {
	var b = [3]int{10, 20, 30}
	valueArray(b)
	fmt.Println(b)
	pointerArray(&b)
	fmt.Println(b)
}

func initArray() {

	var b = [5]bool{}
	fmt.Println(b)

	var d = [5]int64{}
	fmt.Println(d)

	var e = [5]string{}
	fmt.Println(e)

}

func twoArray() {
	var b = [3][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	fmt.Println(b)

}

func runArra() {
	comparePointerAndValue()
	initArray()
	twoArray()
}
