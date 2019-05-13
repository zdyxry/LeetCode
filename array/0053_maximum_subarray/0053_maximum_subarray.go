package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, sumSoFar := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > sumSoFar+nums[i] {
			sumSoFar = nums[i]
		} else {
			sumSoFar += nums[i]
		}
		if sumSoFar > maxSum {
			maxSum = sumSoFar
		}
	}

	return maxSum
}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
