package main

import (
	"fmt"
)

func checkKey() {
	var b = make(map[string]int)
	b = map[string]int{"a": 1, "b": 2}
	value, ok := b["a"]
	fmt.Println(value, ok)
	value, ok = b["d"]
	fmt.Println(value, ok)
}

func deleteKey() {
	var b = make(map[string]int)
	b = map[string]int{"a": 1, "b": 2}
	fmt.Println(b)
	delete(b, "a")
	fmt.Println(b)
}

func main() {
	checkKey()
	deleteKey()
}
