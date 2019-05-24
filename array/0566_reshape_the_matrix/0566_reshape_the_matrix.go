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
		orgY := i / n
		orgX := i % n
		newY := i / c
		newX := i % c
		res[newY][newX] = nums[orgY][orgX]
	}

	return res
}

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrixReshape(nums, 1, 4))
}
