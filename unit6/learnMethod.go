package unit6

import (
	"fmt"
)

type selfInt int

/*func (se selfInt) add(newSelf int) selfInt {
	fmt.Println(selfInt(newSelf))
	se = se + selfInt(newSelf)
	return se
}*/

func (se *selfInt) add(newSelf int) selfInt {
	fmt.Println(selfInt(newSelf))
	*se = *se + selfInt(newSelf)
	return *se
}

func testMethod() {
	var valueSe1 = selfInt(1)
	valueSe2 := valueSe1.add(2)
	fmt.Println(valueSe1, valueSe2)
}

func runMe() {
	testMethod()
}
