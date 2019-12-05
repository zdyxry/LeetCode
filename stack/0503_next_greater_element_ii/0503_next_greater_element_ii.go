package main 

import (
	"fmt"
)

func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))
	stack := make([]int, len(nums))
	top := -1

	for i := range result {
		result[i] = -1
	}

	for i := 0; i<len(nums) * 2 ;i++ {
		num := nums[i % len(nums)]
		for top > -1 && nums[stack[top]] < num {
			result[stack[top]] = num
			top--
		}
		if i < len(nums) {
			top++
			stack[top] = i
		}
	}
	return result
}

func main() {
	nums := []int{1,2,1}
	res := nextGreaterElements(nums)
	fmt.Println(res)
}