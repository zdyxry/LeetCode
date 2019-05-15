package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, n := range nums {
		if m[n] {
			return true
		}
		m[n] = true
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
