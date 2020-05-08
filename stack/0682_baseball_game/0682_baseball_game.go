package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	stack := make([]int, 0)
	for _, s := range ops {
		if s == "C" {
			stack = stack[:len(stack)-1]
		} else if s == "D" {
			stack = append(stack, stack[len(stack)-1]*2)
		} else if s == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		} else {
			temp, _ := strconv.Atoi(s)
			stack = append(stack, temp)
		}
	}

	res := 0
	for _, v := range stack {
		res += v
	}

	return res

}

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	res := calPoints(ops)
	fmt.Println(res)
}
