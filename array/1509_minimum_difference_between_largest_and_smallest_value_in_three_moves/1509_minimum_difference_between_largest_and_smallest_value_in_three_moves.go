package main

import (
	"fmt"
	"math"
	"sort"
)

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	ans := math.MaxInt32
	for i := 0; i <= 3; i++ {
		ans = min(ans, nums[n-4+i]-nums[i])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums := []int{5, 3, 2, 4}
	res := minDifference(nums)
	fmt.Println(res)
}
