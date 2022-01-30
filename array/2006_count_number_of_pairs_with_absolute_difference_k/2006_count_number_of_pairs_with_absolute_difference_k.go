package main

func countKDifference(nums []int, k int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				result += 1
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
