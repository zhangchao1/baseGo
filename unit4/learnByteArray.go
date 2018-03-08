package unit4

import (
	"fmt"
)

func changeStr(c string, s byte) string {
	d := []byte(c)
	d[0] = s
	return string(d)
}

func runByte() {
	fmt.Println(changeStr("123", 'e'))
}
