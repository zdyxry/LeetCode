package main

import "fmt"

func maxSumDivThree(nums []int) int {
	res := 0
	one := 10000
	two := 10000
	for i := 0; i < len(nums); i++ {
		res += nums[i]
		if nums[i]%3 == 1 {
			two = min(two, one+nums[i])
			one = min(one, nums[i])
		}
		if nums[i]%3 == 2 {
			one = min(one, two+nums[i])
			two = min(two, nums[i])
		}
	}
	if res%3 == 0 {
		return res
	} else if res%3 == 1 {
		return res - one
	} else {
		return res - two
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{3, 6, 5, 1, 8}
	res := maxSumDivThree(nums)
	fmt.Println(res)
}
