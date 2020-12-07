package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	ans := 0
	for left < right {
		temp := nums[left] + nums[right]
		if temp < k {
			left++
		} else if temp > k {
			right--
		} else {
			ans++
			left++
			right--
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := maxOperations(nums, 5)
	fmt.Println(res)
}
