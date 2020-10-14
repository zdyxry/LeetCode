package main

import "fmt"

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minSoFar := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if minSoFar > nums[i] {
			minSoFar = nums[i]
		}
	}
	return sum - len(nums)*minSoFar
}

func main() {
	nums := []int{1, 2, 3}
	res := minMoves(nums)
	fmt.Println(res)
}
