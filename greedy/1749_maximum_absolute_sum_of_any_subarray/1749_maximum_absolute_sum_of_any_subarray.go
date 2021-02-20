package main

import "fmt"

func maxAbsoluteSum(nums []int) int {
	pre := 0
	min := 0
	max := 0

	for _, num := range nums {
		pre += num
		if min > pre {
			min = pre
		}
		if max < pre {
			max = pre
		}
	}

	ans := max - min
	if ans < 0 {
		ans = -1 * ans
	}
	return ans
}

func main() {
	nums := []int{1, -3, 2, 3, -4}
	res := maxAbsoluteSum(nums)
	fmt.Println(res)
}
