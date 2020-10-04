package main

import "fmt"

func specialArray(nums []int) int {

	counting := make([]int, 1001)
	for i := range nums {
		counting[nums[i]]++
	}

	for i := len(counting) - 2; i > -1; i-- {
		counting[i] += counting[i+1]
	}

	for i := 0; i <= len(nums); i++ {
		if counting[i] == i {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{3, 5}
	res := specialArray(nums)
	fmt.Println(res)
}
