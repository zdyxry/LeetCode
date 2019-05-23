package main

import (
	"fmt"
)

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums[0]) == 0 || len(nums)*len(nums[0]) != r*c || len(nums) == r && len(nums[0]) == c {
		return nums
	}

	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	m, n := len(nums), len(nums[0])
	for i := 0; i < m*n; i++ {
		org_y := i / n
		org_x := i % n
		new_y := i / c
		new_x := i % c
		res[new_y][new_x] = nums[org_y][org_x]
	}

	return res
}

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrixReshape(nums, 1, 4))
}
