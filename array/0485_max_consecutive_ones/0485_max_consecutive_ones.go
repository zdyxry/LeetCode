package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	max, count := 0, 0
	for _, v := range nums {
		switch v {
		case 1:
			count++
			if count > max {
				max = count
			}
		case 0:
			count = 0
		}
	}

	return max
}

func main() {
	nums := []int{1, 0, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
