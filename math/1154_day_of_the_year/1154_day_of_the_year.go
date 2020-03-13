package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	order := 0
	ds := strings.Split(date, "-")
	if len(ds) != 3 {
		return order
	}
	days := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, _ := strconv.Atoi(ds[0])
	m, _ := strconv.Atoi(ds[1])
	d, _ := strconv.Atoi(ds[2])
	order += d
	for i := 0; i < m-1; i++ {
		order += days[i]
	}
	if m > 2 && (y%100 == 0 && y%400 == 0) || (y%100 != 0 && y%4 == 0) {
		order++
	}
	return order
}

func main() {
	res := dayOfYear("2019-01-09")
	fmt.Println(res)
}
