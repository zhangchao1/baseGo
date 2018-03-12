package baseStruct

import (
	"fmt"
	"reflect"
)

type father struct {
	f int64
}

type child struct {
	s int64
	father
}

type fac struct {
	id   int64
	name string
}

type tagStruct struct {
	id   int64  "orm id"
	name string "string 123"
}

func runNestStruct() {
	c := new(child)
	c.f = 1
	c.s = 2
	fmt.Println(c)
}

func NewFac(id int64, name string) *fac {
	return &fac{id, name}
}

func structFiled() {
	ts := tagStruct{id: 1, name: "123"}
	for i := 0; i < 2; i++ {
		refTag(ts, i)
	}
}

func refTag(tt tagStruct, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func runs() {
	newFac := NewFac(1, "chao")
	fmt.Println(newFac)
	structFiled()
}
