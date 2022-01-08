package main

func buildArray(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		result = append(result, nums[num])
	}
	return result
}
