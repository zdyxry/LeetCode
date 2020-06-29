package main

import "fmt"

func longestSubarray(nums []int) int {
	l, r := 0, 0
	zeroIdx := -1

	maxOnes := 0
	for r < len(nums) {
		if nums[r] == 0 {
			l = zeroIdx + 1
			zeroIdx = r
		}
		maxOnes = getMax(maxOnes, r-l)
		r++
	}

	return maxOnes
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	nums := []int{1, 1, 0, 0, 1}
	res := longestSubarray(nums)
	fmt.Println(res)
}
