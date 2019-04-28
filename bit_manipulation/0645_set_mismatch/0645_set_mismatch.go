package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	dup := 0
	for i := 0; i < len(nums); i++ {
		n := abs(nums[i])

		if nums[n-1] < 0 {
			dup = n
		} else {
			nums[n-1] = -nums[n-1]
		}
	}

	mis := 0
	for i, v := range nums {
		if v > 0 {
			mis = i + 1
			break
		}
	}

	return []int{dup, mis}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	nums := []int{1, 2, 2, 4}
	fmt.Println(findErrorNums(nums))
}
