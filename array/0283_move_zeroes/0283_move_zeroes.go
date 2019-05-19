package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	pos := 0
	for idx, _ := range nums {
		if nums[idx] != 0 {
			nums[idx], nums[pos] = nums[pos], nums[idx]
			pos += 1
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
