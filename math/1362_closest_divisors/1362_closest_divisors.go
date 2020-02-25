package main

import (
	"fmt"
	"math"
)

func closestDivisors(num int) []int {
	v1 := helper(num + 1)
	v2 := helper(num + 2)
	if v1[1]-v1[0] < v2[1]-v2[0] {
		return v1
	} else {
		return v2
	}
}

func helper(num int) []int {
	root := int(math.Sqrt(float64(num)))
	for i := root; i >= 1; i-- {
		if num%i == 0 {
			return []int{i, num / i}
		}
	}
	return []int{}
}

func main() {
	num := 8
	res := closestDivisors(num)
	fmt.Println(res)
}
