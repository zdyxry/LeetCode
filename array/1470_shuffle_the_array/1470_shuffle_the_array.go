package main

import "fmt"

func shuffle(nums []int, n int) []int {
	res := []int{}
	for i := 0; i < n; i++ {
		res = append(res, nums[i])
		if i+n < 2*n {
			res = append(res, nums[i+n])
		}
	}
	return res
}

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	res := shuffle(nums, 3)
	fmt.Println(res)
}
