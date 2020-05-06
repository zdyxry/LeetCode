package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	if k == 0 {
		return true
	}
	var prevIdx int = -k - 1
	for idx, num := range nums {
		if num == 1 {
			if idx-prevIdx <= k {
				return false
			} else {
				prevIdx = idx
			}
		}
	}
	return true
}

func main() {
	nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
	res := kLengthApart(nums, 2)
	fmt.Println(res)
}
