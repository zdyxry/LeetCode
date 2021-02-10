package main

import "fmt"

func check(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			cnt++
			if cnt > 1 || nums[0] < nums[len(nums)-1] {
				return false
			}
		}

	}
	return true
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	res := check(nums)
	fmt.Println(res)
}
