package main

import (
	"fmt"
	"math"
)

func mctFromLeafValues(arr []int) int {
	stack := []int{math.MaxInt32}

	res := 0

	for _, v := range arr {
		for stack[len(stack)-1] <= v {
			num := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = res + num*int(math.Min(float64(v), float64(stack[len(stack)-1])))
		}
		stack = append(stack, v)
	}

	for len(stack) > 2 {
		num := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = res + num*stack[len(stack)-1]
	}
	return res
}

func main() {
	arr := []int{6, 2, 4, 8}
	res := mctFromLeafValues(arr)
	fmt.Println(res)
}
