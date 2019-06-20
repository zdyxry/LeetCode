package main

import (
	"fmt"
)

func findLHS(nums []int) int {
	r := make(map[int]int, len(nums))
	for _, n := range nums {
		r[n]++
	}

	max := 0
	for n, c1 := range r {
		c2, ok := r[n+1]
		if ok {
			t := c1 + c2
			if max < t {
				max = t
			}
		}
	}

	return max
}
func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(nums))
}
