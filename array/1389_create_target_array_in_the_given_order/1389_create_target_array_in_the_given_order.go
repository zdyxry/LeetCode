package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(index))
	for i := range res {
		res[i] = -1
	}

	for i, indexI := range index {
		if res[indexI] != -1 {
			copy(res[indexI+1:], res[indexI:])
		}
		res[indexI] = nums[i]
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	res := createTargetArray(nums, index)
	fmt.Println(res)
}
