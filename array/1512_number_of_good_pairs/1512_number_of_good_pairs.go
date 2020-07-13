package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	var index1, index2 int
	var r int
	for index1 < len(nums) && index2 < len(nums) {
		if index1 < index2 {
			if nums[index1] == nums[index2] {
				r++
			}
		}
		index2++
		if index2 == len(nums) {
			index1++
			index2 = 0
		}
		if index1 == len(nums)-1 {
			break
		}
	}
	return r
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	res := numIdenticalPairs(nums)
	fmt.Println(res)
}
