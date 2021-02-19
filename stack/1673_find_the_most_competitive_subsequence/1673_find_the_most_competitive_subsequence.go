package main

import "fmt"

func mostCompetitive(nums []int, k int) []int {
	res := []int{}
	k = len(nums) - k
	for i := range nums {
		for len(res) > 0 && k > 0 && res[len(res)-1] > nums[i] {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, nums[i])
	}
	return res[:len(res)-k]
}

func main() {
	nums := []int{3, 5, 2, 6}
	k := 2
	res := mostCompetitive(nums, k)
	fmt.Println(res)
}
