package main

import (
	"fmt"
	"sort"
)

func minMoves2(nums []int) int {
	lenth := len(nums)
	sort.Ints(nums)
	num := nums[lenth/2]
	res := 0
	for i := 0; i < lenth; i++ {
		res += getAbs(num, nums[i])
	}
	return res
}
func getAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	nums := []int{1, 2, 3}
	res := minMoves2(nums)
	fmt.Println(res)
}
