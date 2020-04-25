package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	sign := ""
	if num < 0 {
		num *= -1
		sign = "-"
	}
	base7 := ""
	for num != 0 {
		base7 = strconv.Itoa(num%7) + base7
		num /= 7
	}
	return sign + base7
}

func main() {
	num := 100
	res := convertToBase7(num)
	fmt.Println(res)
}
