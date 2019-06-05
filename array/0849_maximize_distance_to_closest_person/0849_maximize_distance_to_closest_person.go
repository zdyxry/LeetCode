package main

import (
	"fmt"
)

func maxDistToClosest(seats []int) int {
	start := 0
	max := 0
	var i int
	for ; i < len(seats); i++ {
		if seats[i] == 1 {
			tmp := i - start
			if start != 0 {
				tmp = (tmp + 1) / 2
			}
			if max < tmp {
				max = tmp
			}
			start = i + 1
		}
	}
	if max < i-start {
		max = i - start
	}

	return max
}

func main() {
	seats := []int{1, 0, 0, 0, 1, 0, 1}
	fmt.Println(maxDistToClosest(seats))
}
