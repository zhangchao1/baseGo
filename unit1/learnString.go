package uint1

import (
	"fmt"
	"strconv"
	"strings"
)

func convertStringInt(a string) (b int, err error) {
	b, err = strconv.Atoi(a)
	fmt.Println(b, err)
	return b, err
}

func convertStringToFloat(a string) (b float64, err error) {
	b, err = strconv.ParseFloat(a, 64)
	fmt.Println(b, err)
	return b, err
}

func convertIntToString(a int) (b string) {
	b = strconv.Itoa(a)
	fmt.Println(b)
	return b
}

func convertFloatToString(a float64) (b string) {
	b = strconv.FormatFloat(a, 'f', 4, 64)
	fmt.Println(b)
	return b
}

func runString() {
	convertStringInt("666")
	convertStringToFloat("666.89")
	convertIntToString(123)
	convertFloatToString(1.234)
	/*fenge
	string的相关的操作
	*/

	b := []string{"a", "b", "c"}
	fmt.Println(strings.Contains("abc", "a"))        //字符串包含
	fmt.Println(strings.Index("abc", "b"))           //找出字符串出现的位置
	fmt.Println(strings.Replace("abc", "b", "e", 1)) //替换对应位置的字符串
	fmt.Println(strings.Count("aababab", "a"))       //统计字符串出现的次数
	fmt.Println(strings.Split("a,b,c", ","))         //字符串分割
	fmt.Println(strings.Join(b, ","))                //字符串拼接
	fmt.Println(strings.ToUpper("abc"))              //转成大写
	fmt.Println(strings.ToLower("ABC"))              //转成小写
}
