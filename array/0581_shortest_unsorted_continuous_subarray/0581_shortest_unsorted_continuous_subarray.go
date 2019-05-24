package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := 0, -1
	min, max := nums[n-1], nums[0]

	for i := 1; i < n; i++ {
		if max <= nums[i] {
			max = nums[i]
		} else {
			right = i
		}

		j := n - i - 1
		if min >= nums[j] {
			min = nums[j]
		} else {
			left = j
		}
		fmt.Println(i, left, right, max, min)
	}

	return right - left + 1
}
func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(nums))
}
