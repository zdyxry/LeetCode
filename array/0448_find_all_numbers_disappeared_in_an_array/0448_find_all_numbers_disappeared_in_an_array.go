package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			// fmt.Println(nums)
		}
	}
	// fmt.Println(nums)

	res := make([]int, 0, len(nums))

	for i, n := range nums {
		if n != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
