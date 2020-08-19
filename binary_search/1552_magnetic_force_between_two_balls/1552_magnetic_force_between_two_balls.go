package main

import (
	"fmt"
	"sort"
)

func search(dist, m int, position []int) bool {
	curr, res := 0, 1
	for i := 0; i < len(position)-1; i++ {
		curr += position[i+1] - position[i]
		if curr >= dist {
			res += 1
			curr = 0
		}
	}
	return res >= m
}

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	l, r := 1, position[n-1]-position[0]

	for l < r {
		mid := (l + r + 1) / 2
		if search(mid, m, position) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	position := []int{1, 2, 3, 4, 7}
	m := 3
	res := maxDistance(position, m)
	fmt.Println(res)
}
