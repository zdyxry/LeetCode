package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 2, 3, 4, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}
