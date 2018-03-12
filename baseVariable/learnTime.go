package baseVariable

import (
	"fmt"
	"time"
)

//类似于Y-m-d H:i:s的形式
func formatTimeToString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func timeNow() (int, int) {
	return time.Now().Day(), time.Now().Year()
}

//字符串转int类型的时间
func timeConvertStringToInt(a string) int64 {
	tempTime, _ := time.Parse("2006-01-02 15:04:05", a)
	return tempTime.Unix()
}

//int类型转string类型
func timeConvertIntToString() {
	str_time := time.Unix(1989058332, 0).Format("2006-01-02 15:04:05")
	fmt.Println(str_time)
}

func compareTime(a string, b string) int {
	a1, _ := time.Parse("2006-01-02 15:04:05", a)
	b1, _ := time.Parse("2006-01-02 15:04:05", b)
	if a1.After(b1) {
		return 1
	} else if a1.Before(b1) {
		return -1
	} else {
		return 0
	}
}

func addDays(n int, format string) string {
	tempTime := time.Now()
	addTime := tempTime.AddDate(0, 0, n)
	return addTime.Format(format)
}

func runTime() {
	fmt.Println(formatTimeToString())
	fmt.Println(timeNow())
	fmt.Println(timeConvertStringToInt("2018-03-01 01:35:00"))
	fmt.Println(time.Now().Unix())
	timeConvertIntToString()

	//两个时间的比较
	fmt.Println(compareTime("2018-09-01 00:00:00", "2018-07-09 00:00:00"))
	fmt.Println(compareTime("2018-01-01 00:00:00", "2018-07-09 00:00:00"))
	fmt.Println(compareTime("2018-07-09 00:00:00", "2018-07-09 00:00:00"))

	fmt.Println(addDays(1, "20060102"))  //明天的时间
	fmt.Println(addDays(-1, "20060102")) //昨天的时间
}
