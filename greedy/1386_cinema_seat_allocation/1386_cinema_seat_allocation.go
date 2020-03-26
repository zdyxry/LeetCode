package main

import "fmt"

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	blocked := make(map[int]int)
	for _, seat := range reservedSeats {
		row, col := seat[0], seat[1]-1
		blocked[row] |= 1 << col
	}

	count := 2 * n
	for _, row := range blocked {
		left := (row & 0b11110) != 0
		center := (row & 0b1111000) != 0
		right := (row & 0b111100000) != 0
		if left && center && right {
			count -= 2
		} else if !(!left && !right) {
			count--
		}
	}
	return count
}

func main() {
	n := 3
	reservedSeats := [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}
	res := maxNumberOfFamilies(n, reservedSeats)
	fmt.Println(res)
}
