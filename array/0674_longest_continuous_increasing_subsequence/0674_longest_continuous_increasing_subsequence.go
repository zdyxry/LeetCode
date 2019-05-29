package main

import (
	"fmt"
	"math"
)

func findLengthOfLCIS(nums []int) int {
	res, cur := 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] < nums[i] {
			cur++
		} else {
			res = int(math.Max(float64(res), float64(cur)))
			cur = 1
		}
	}
	return int(math.Max(float64(res), float64(cur)))
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums))
}
