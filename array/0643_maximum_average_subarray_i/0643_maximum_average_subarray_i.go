package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	total := 0
	for i := 0; i < k; i++ {
		total += nums[i]
	}

	result := total

	for i := k; i < len(nums); i++ {
		total = total + nums[i] - nums[i-k]
		if result < total {
			result = total
		}
	}
	return float64(result) / float64(k)
}
func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}
