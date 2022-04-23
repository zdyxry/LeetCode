package main

func findClosestNumber(nums []int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if abs(nums[i]) < abs(res) || abs(nums[i]) == abs(res) && nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
