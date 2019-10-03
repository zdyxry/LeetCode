package main

import "fmt"

func removeDuplicates(nums []int) int {
	size := len(nums)

	i := 2

	for j :=i;j < size; j++ {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

func main() {
	nums := []int{1,1,1,2,2,3}
	res := removeDuplicates(nums)
	fmt.Println(res)
}