package main

import "fmt"

func dayOfTheWeek(day int, month int, year int) string {
	startDate := 3 // 周一到周日分别用0-6表示
	dataMap := map[int]string{
		0: "Monday",
		1: "Tuesday",
		2: "Wednesday",
		3: "Thursday",
		4: "Friday",
		5: "Saturday",
		6: "Sunday",
	}

	// 每个月的天数
	mouthDays := [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	// 记录一共过去了几天
	totalDays := 0
	// 计算给出的今年1.1之前的天数
	for i := 1971; i < year; i++ {

		for j := 1; j <= 12; j++ {
			totalDays += mouthDays[j-1]
		}

		// 如果是闰年， 天数+1
		if i%400 == 0 || (i%4 == 0 && i%100 != 0) {
			totalDays += 1
		}
	}

	// 计算今年当前月之前的天数,
	for i := 1; i < month; i++ {
		totalDays += mouthDays[i-1]
		if i == 2 && (year%400 == 0 || (year%4 == 0 && year%100 != 0)) {
			totalDays += 1
		}
	}
	totalDays += day

	return dataMap[(totalDays+startDate)%7]
}

func main() {
	fmt.Println(dayOfTheWeek(15, 10, 1986))
}
