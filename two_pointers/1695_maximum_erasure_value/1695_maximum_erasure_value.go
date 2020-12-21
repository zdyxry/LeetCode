package main

import "fmt"

func maximumUniqueSubarray(nums []int) int {

	count := make([]int, 10001)
	left := 0
	res := nums[0]
	sum := nums[0]
	count[nums[0]]++
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		count[nums[i]]++
		for count[nums[i]] > 1 {
			count[nums[left]]--
			sum -= nums[left]
			left++
		}
		res = max(res, sum)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{4, 2, 4, 5, 6}
	res := maximumUniqueSubarray(nums)
	fmt.Println(res)
}
