package main

func countPairs(nums []int, k int) int {
	var res int
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				res++
			}
		}
	}
	return res
}
