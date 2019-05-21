package main

import (
	"fmt"
)

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	record := make(map[int]int)
	for _, num := range nums {
		record[num]++
	}

	ans := 0

	if k == 0 {
		for _, count := range record {
			if count > 1 {
				ans++
			}
		}
		return ans
	}

	for n := range record {
		if record[n-k] > 0 {
			ans++
		}
	}

	return ans
}
func main() {
	nums := []int{3, 1, 4, 1, 5}
	k := 2
	fmt.Println(findPairs(nums, k))
}
