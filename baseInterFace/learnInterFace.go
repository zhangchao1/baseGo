package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

type dayArray struct {
	data []*day
}
type day struct {
	num       int
	shortName string
	longName  string
}

type Appender interface {
	Append(int)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
	fmt.Println(a)
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	/*var list1 List
	fmt.Println(LongEnough(list1))
	var lis = new(List)
	CountInto(lis, 0, 10)*/
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	fmt.Println(a)
}
