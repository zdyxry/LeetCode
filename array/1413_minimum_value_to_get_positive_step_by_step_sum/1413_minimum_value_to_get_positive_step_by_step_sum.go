package main

import "fmt"

func minStartValue(nums []int) int {
	minSum := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		minSum = min(minSum, sum)
	}
	return max(1, 1-minSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{-3, 2, -3, 4, 2}
	res := minStartValue(nums)
	fmt.Println(res)
}
