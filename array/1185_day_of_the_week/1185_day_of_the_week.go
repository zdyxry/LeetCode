package main

import "fmt"

func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

func getDay(year int, mon int, day int) (res int) {
	months := []int{-1, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := 1970; i < year; i++ {
		res += 365
		if isLeapYear(i) {
			res++
		}
	}
	for i := 1; i < mon; i++ {
		res += months[i]
		if i == 2 && isLeapYear(year) {
			res++
		}
	}
	res += day
	return res
}

func dayOfTheWeek(day int, month int, year int) string {
	weekMap := map[int]string{
		1: "Thursday",
		2: "Friday",
		3: "Saturday",
		4: "Sunday",
		5: "Monday",
		6: "Tuesday",
		0: "Wednesday",
	}
	d := getDay(year, month, day)
	return weekMap[d%7]
}

func main() {
	day := 31
	month := 8
	year := 2019
	res := dayOfTheWeek(day, month, year)
	fmt.Println(res)
}
