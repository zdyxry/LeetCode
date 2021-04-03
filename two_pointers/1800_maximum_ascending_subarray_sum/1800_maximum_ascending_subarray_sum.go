package main

import "fmt"

func maxAscendingSum(nums []int) int {
	max, temp := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp += nums[i]
		} else {
			temp = nums[i]
		}
		max = Compare(max, temp)
	}
	return max
}

func Compare(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{10, 20, 30, 5, 10, 50}
	res := maxAscendingSum(nums)
	fmt.Println(res)
}
