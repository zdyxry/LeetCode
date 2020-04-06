package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	var (
		i, sum, tmp int
		tmpSlice    []int
		length      = len(nums)
	)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i = 0; i < length; i++ {
		sum += nums[i]
	}
	for i = 0; i < length; i++ {
		tmp += nums[i]
		tmpSlice = append(tmpSlice, nums[i])
		if tmp > sum-tmp {
			return tmpSlice
		}
	}
	return tmpSlice
}

func main() {
	nums := []int{4, 3, 10, 9, 8}
	res := minSubsequence(nums)
	fmt.Println(res)
}
