package main

import (
	"fmt"
)

func checkPossibility(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 || len(nums) == 2 {
		return true
	}

	found := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if found {
				return false
			}
			if i-2 >= 0 && nums[i] < nums[i-2] {
				nums[i] = nums[i-1]
			}
			found = true
		}
	}

	return true
}
func main() {
	nums := []int{4, 2, 3}
	fmt.Println(checkPossibility(nums))
}
