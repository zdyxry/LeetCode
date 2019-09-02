package main 

import (
	"fmt"
)

func search(nums []int, target int) int {
	l, r  := 0, len(nums) - 1

	for l <= r{
		m := (l + r) / 2
		switch {
		case nums[m] < target:
			l = m + 1
		case target < nums[m]:
			r = m - 1
		default:
			return m
		}
	}

	return -1
}

func main() {
	nums := []int{-1,0,3,5,9,12}
	target := 9
	res := search(nums, target)
	fmt.Println(res)
}