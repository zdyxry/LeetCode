package main

import (
	"fmt"
)

func nextGreaterElement(findNums []int, nums []int) []int {
	indexOf := make(map[int]int)
	for i, n := range nums {
		indexOf[n] = i
	}

	res := make([]int, len(findNums))
	for i, n := range findNums {
		res[i] = -1
		for j := indexOf[n] +1; j < len(nums); j++ {
			if n < nums[j] {
				res[i] = nums[j]
				break
			}
		}
	}

	return res
}

func main() {
	findNums := []int{4,1,2}
	nums := []int{1,3,4,2}
	res := nextGreaterElement(findNums, nums)
	fmt.Println(res)
}