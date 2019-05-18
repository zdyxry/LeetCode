package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	xor := 0

	for i, n := range nums {
		xor ^= i ^ n
	}

	return xor ^ len(nums)
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}
